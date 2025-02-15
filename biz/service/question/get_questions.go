package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/tongque0/edupal/biz/dal/cache"
	"github.com/tongque0/edupal/biz/dal/model"
	"github.com/tongque0/edupal/biz/dal/mysql"
	question "github.com/tongque0/edupal/hertz_gen/question"
)

type GetQuestionsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetQuestionsService(Context context.Context, RequestContext *app.RequestContext) *GetQuestionsService {
	return &GetQuestionsService{RequestContext: RequestContext, Context: Context}
}
func (h *GetQuestionsService) Run(req *question.GetQuestionReq) (resp *question.GetQuestionResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// 查询库里所有符合条件的记录
	var questions []model.Question

	// 构建查询条件
	query := mysql.DB.Model(&model.Question{})
	if req.Ownerid != "" {
		query = query.Where("ower_id = ?", req.Ownerid)
	}
	if req.Quesbankid != "" {
		query = query.Where("ques_bank_id = ?", req.Quesbankid)
	}
	if req.Quesid != "" {
		query = query.Where("ques_id = ?", req.Quesid)
	}
	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}
	if req.Subject != "" {
		query = query.Where("subject = ?", req.Subject)
	}
	if req.Genid != "" {
		query = query.Where("gen_id = ?", req.Genid)
	}

	result := query.Find(&questions)
	if result.Error != nil {
		hlog.CtxErrorf(h.Context, "查询题目失败: %v", result.Error)
		return nil, result.Error
	}

	// 打印查询到的题目
	hlog.CtxInfof(h.Context, "查询到的题目数量: %d", len(questions))

	// 构建响应
	resp = &question.GetQuestionResp{
		Questions: make([]*question.Question, 0, len(questions)),
		HasMore:   false,
	}

	for _, ques := range questions {
		resp.Questions = append(resp.Questions, &question.Question{
			Type:       ques.Type,
			OwnerId:    ques.OwerID,
			QuesId:     ques.QuesID,
			QuesBankId: ques.QuesBankID,
			Title:      ques.Title,
			Question:   ques.Question,
			Answer:     ques.Answer,
			Tip:        ques.Tip,
			Parse:      ques.Parse,
		})
	}

	// 检查缓存中是否存在genid
	if req.Genid != "" {
		if _, found := cache.DB.Get(req.Genid); found {
			resp.HasMore = true
		}
	}

	return resp, nil
}
