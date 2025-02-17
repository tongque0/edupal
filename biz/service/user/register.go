package service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/duke-git/lancet/v2/random"
	"github.com/tongque0/edupal/biz/dal/model"
	"github.com/tongque0/edupal/biz/dal/mysql"
	user "github.com/tongque0/edupal/hertz_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}

	usr := model.User{}
	//查询用户是否存在
	err = mysql.DB.Where("username = ?", req.Username).First(&usr).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.New("服务器错误")
		}
	}
	//用户已存在
	if usr.ID != 0 {
		return nil, errors.New("用户已存在")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码哈希失败")
	}
	userid, err := random.UUIdV4()
	if err != nil {
		hlog.CtxErrorf(h.Context, "生成 UUID 失败: %v", err)
		return
	}
	// 创建新用户
	newUser := &model.User{
		UserID:         userid,
		Username:       req.Username,
		PasswordHashed: string(hashedPassword),
		Role:           "user",
	}

	err = mysql.DB.Create(newUser).Error
	if err != nil {
		return nil, errors.New("创建用户失败")
	}

	// todo edit your code
	return &user.RegisterResp{
		UserId: newUser.UserID,
	}, nil
}
