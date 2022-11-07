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

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 从数据库中查询当前用户
	user := new(model.UserBasic)
	get, err := model.Engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)

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
