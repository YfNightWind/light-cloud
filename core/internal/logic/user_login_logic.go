package logic

import (
	"context"
	"light-cloud/src/core/define"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	resp = new(types.LoginResponse)

	// 从数据库中查询当前用户
	user := new(model.UserInfo)
	if len(req.Name) < 6 {
		resp.Msg = "用户名长度不能小于6位"
		return
	}
	if len(req.Password) < 6 {
		resp.Msg = "密码长度不能小于6位"
		return
	}

	get, err := l.svcCtx.SQL.Where("name = ? AND password = ? ", req.Name, helper.Md5(req.Password)).Get(user)

	if err != nil {
		return nil, err
	}

	if !get {
		resp.Msg = "用户名或密码错误"
		return
	}

	// 生成普通 token1
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		resp.Msg = "生成token失败"
		return
	}

	// 生成用于刷新 token1 的 token2
	// 当 token1 失效后，使用 token2 生成新 token1
	refreshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		resp.Msg = "生成token失败"
		return
	}

	resp.Token = token
	resp.RefreshToken = refreshToken
	resp.Msg = "用户登录成功"

	return
}
