package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/tongque0/edupal/biz/dal/model"
	"github.com/tongque0/edupal/biz/dal/mysql"
	"github.com/tongque0/edupal/conf"
	question "github.com/tongque0/edupal/hertz_gen/question"
)

type CheckAnswerService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckAnswerService(Context context.Context, RequestContext *app.RequestContext) *CheckAnswerService {
	return &CheckAnswerService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckAnswerService) Run(req *question.CheckAnswerReq) (resp *question.CheckResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	quesid := req.GetQuesid()
	answer := req.GetAnswer()
	// 查询题目
	var ques model.Question
	result := mysql.DB.Where("ques_id = ?", quesid).First(&ques)
	if result.Error != nil {
		hlog.CtxErrorf(h.Context, "查询题目失败: %v", result.Error)
		return nil, result.Error
	}
	checkResult, err := checkAnswer(h.Context, ques, answer)
	if err != nil {
		hlog.CtxErrorf(h.Context, "检查答案失败: %v", err)
		return nil, err
	}

	return &question.CheckResp{
		Result:    checkResult.Result,
		ResultMsg: checkResult.ResultMsg,
	}, nil
}

type Check struct {
	Result    string `json:"result"`
	ResultMsg string `json:"result_msg"`
}

func checkAnswer(ctx context.Context, question model.Question, answer string) (Check, error) {
	config := openai.DefaultConfig(conf.GetConf().AI.APIKey)
	config.BaseURL = conf.GetConf().AI.BaseURL
	client := openai.NewClientWithConfig(config)

	var result Check
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		hlog.CtxErrorf(ctx, "生成结构化输出Schema错误: %v", err)
	}

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o20241120,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `你是一位资深的中国初高中教师，拥有多年的教学经验，擅长根据不同年级、科目和难度判定学生的答案。在判定学生答案时，如果题库答案不完全正确或者存在争议，请根据自己的判断提供合理的判定和反馈。每个问题的反馈应包括：
				1. 答案是否正确（“正确”或“错误”）。
				2. 反馈说明，简短地解释为什么答案是正确或错误。
				3. 如果答案错误，给出具体的改正建议或提示。
				4. 如果答案正确，提供一个详细的解析，帮助学生理解正确的解题思路和过程。
				5. 如果题库的答案存在问题或不完全准确，请结合你的判断和经验，给出更合理的评判。如果你认为题库的答案存在明显错误，则提供正确的答案或解析。确保最终的判定具有教育意义。`,
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: fmt.Sprintf(`请对以下问题进行判定：
				问题：%s 题库答案：%s 学生答案：%s
				请判断学生的答案是否正确，并给出反馈。`,
					question.Question, question.Answer, answer),
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:        "check_answer",
				Description: "根据学生的答案进行判题，并提供详细反馈。",
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
	return result, nil
}
