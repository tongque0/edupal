package service

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/paseto"
	"github.com/tongque0/edupal/biz/dal/model"
	"github.com/tongque0/edupal/biz/dal/mysql"
	"github.com/tongque0/edupal/conf"
	user "github.com/tongque0/edupal/hertz_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}

	//查询用户
	usr := model.User{}
	err = mysql.DB.Where("username = ?", req.Username).First(&usr).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, errors.New("服务器错误")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	if usr.Role == "" {
		usr.Role = "user"
	}
	now := time.Now()
	genTokenFunc := paseto.DefaultGenTokenFunc()
	token, err := genTokenFunc(&paseto.StandardClaims{
		Issuer:    conf.GetConf().Hertz.JWTSecret,
		ExpiredAt: now.Add(24 * time.Hour),
		NotBefore: now,
		IssuedAt:  now,
	}, utils.H{
		"username": usr.Username,
		"role":     usr.Role,
	}, nil)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}

	return &user.LoginResp{
		UserId: usr.UserID,
		Token:  token,
	}, nil
}
