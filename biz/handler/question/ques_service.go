package question

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	service "github.com/tongque0/edupal/biz/service/question"
	"github.com/tongque0/edupal/biz/utils"
	question "github.com/tongque0/edupal/hertz_gen/question"
)

// GenQuestions .
// @router /question/gen [POST]
func GenQuestions(ctx context.Context, c *app.RequestContext) {
	var err error
	var req question.GenQuestionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &question.GenQuestionResp{}
	resp, err = service.NewGenQuestionsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetQuestions .
// @router /question/get [GET]
func GetQuestions(ctx context.Context, c *app.RequestContext) {
	var err error
	var req question.GetQuestionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetQuestionsService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CheckAnswer .
// @router /question/check [POST]
func CheckAnswer(ctx context.Context, c *app.RequestContext) {
	var err error
	var req question.CheckAnswerReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckAnswerService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
