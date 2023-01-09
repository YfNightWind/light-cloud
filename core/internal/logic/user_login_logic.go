package logic

import (
	"context"
	"errors"
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
	// 从数据库中查询当前用户
	user := new(model.UserInfo)
	get, err := model.Engine.Where("name = ? AND password = ? ", req.Name, helper.Md5(req.Password)).Get(user)

	if err != nil {
		return nil, err
	}

	if !get {
		return nil, errors.New("用户名或密码错误！")
	}

	// 返回token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}

	resp = new(types.LoginResponse)
	resp.Token = token
	return
}
