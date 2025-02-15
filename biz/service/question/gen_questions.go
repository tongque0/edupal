package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/duke-git/lancet/v2/random"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/tongque0/edupal/biz/dal/cache"
	"github.com/tongque0/edupal/biz/dal/model"
	"github.com/tongque0/edupal/biz/dal/mysql"
	"github.com/tongque0/edupal/conf"
	question "github.com/tongque0/edupal/hertz_gen/question"
)

type GenQuestionsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGenQuestionsService(Context context.Context, RequestContext *app.RequestContext) *GenQuestionsService {
	return &GenQuestionsService{RequestContext: RequestContext, Context: Context}
}

func (h *GenQuestionsService) Run(req *question.GenQuestionReq) (resp *question.GenQuestionResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// 生成唯一的genid
	genid, err := random.UUIdV4()
	if err != nil {
		hlog.CtxErrorf(h.Context, "生成 UUID 失败: %v", err)
		return nil, err
	}

	// 异步处理生成题目的任务
	go h.generateQuestions(req, genid)

	// 立即返回 genid
	return &question.GenQuestionResp{
		Genid: genid,
	}, nil
}

func (h *GenQuestionsService) generateQuestions(req *question.GenQuestionReq, genid string) {
	//设置3分钟超时的genid，用于跟踪问题生成任务
	cache.DB.Set(genid, "问题生成中，该缓存将于180s后失效", 180*time.Second)
	subject := req.GetSubject()
	class := req.GetClass()
	diff := req.GetDifficulty()
	knowledge := req.GetKnowledge()
	// user := h.RequestContext.Keys["username"].(string)
	user := "tongque"
	var wg sync.WaitGroup

	// 创建一个带有上下文的子上下文，用于控制 Goroutines
	ctx, cancel := context.WithCancel(h.Context)
	defer cancel()

	// 生成题目
	generate := func(num int32, questype string) {
		for i := 0; i < int(num); i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				question, err := genQuestions(ctx, subject, class, diff, knowledge, questype)
				if err != nil {
					hlog.CtxErrorf(ctx, "生成%s(任务批次:%s第%d个)失败: %v", questype, genid, i, err)
					return
				}
				quesid, err := random.UUIdV4()
				if err != nil {
					hlog.CtxErrorf(h.Context, "生成 UUID 失败: %v", err)
					return
				}
				mysql.DB.Create(&model.Question{
					Type:       questype,
					OwerID:     user,
					OwerName:   user,
					QuesID:     quesid,
					GenID:      genid,
					Title:      question.Title,
					Question:   question.Question,
					Answer:     question.Answer,
					Tip:        question.Tip,
					Parse:      question.Parse,
					Subject:    subject,
					Grade:      class,
					Difficulty: diff,
				})
			}(i)
		}
	}

	generate(req.GetChoiceNum(), "选择题")

	generate(req.GetCalcNum(), "计算题")

	generate(req.GetShortAnsNum(), "简答题")

	// 等待所有任务完成
	wg.Wait()
	cache.DB.Delete(genid)
	hlog.CtxInfof(ctx, "所有题目生成任务已完成")
}

type Result struct {
	Title    string `json:"title"`    //标题
	Question string `json:"question"` //题目
	Answer   string `json:"answer"`   //答案
	Tip      string `json:"tip"`      //提示
	Parse    string `json:"parse"`    //解析
}

// genQuestions 根据学科、班级、难度、知识点和题型生成题目
func genQuestions(ctx context.Context, subject, class, diff, knowledge, questype string) (Result, error) {

	config := openai.DefaultConfig(conf.GetConf().AI.APIKey)
	config.BaseURL = conf.GetConf().AI.BaseURL
	client := openai.NewClientWithConfig(config)

	var result Result
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		hlog.CtxErrorf(ctx, "生成结构化输出Schema错误: %v", err)
	}

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o20241120,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `你是一位资深的中国初高中教师，拥有多年的教学经验，擅长根据不同年级、科目和难度设计练习题和考试题。你将根据以下要求为学生生成题目，每道题目包含以下内容：
				1. 题目内容
				2. 答案
				3. 提示（通俗简短的）
				4. 解析（详细讲解题目的解答过程）
				请确保题目具有适当的难度，符合学生的学习进度，并且答案正确、解析清晰。`,
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: fmt.Sprintf(`请根据以下要求生成一道题目：
				科目：%s
				年级：%s
				难度：%s
				知识点：%s
				题型：%s
				`,
					subject, class, diff, knowledge, questype),
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:        "create_question",
				Description: "根据用户要求生成一道练习题，并按以下格式输出：标题，题目内容，答案，提示（如果有），解析。",
				Schema:      schema,
				Strict:      true,
			},
		},
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "创建生成对话失败: %v", err)
	}
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		hlog.CtxErrorf(ctx, "解析生成数据失败: %v", err)
	}

	// fmt.Println("生成题目:")
	// fmt.Printf("Question: %s\n", result.Question)
	// fmt.Printf("Answer: %s\n", result.Answer)
	// fmt.Printf("Tip: %s\n", result.Tip)
	// fmt.Printf("Parse: %s\n", result.Parse)
	// fmt.Println()

	hlog.CtxInfof(ctx, "生成题目: subject=%s, class=%s, diff=%s, knowledge=%s, questype=%d", subject, class, diff, knowledge, questype)
	return result, nil
}
