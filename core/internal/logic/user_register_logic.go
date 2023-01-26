package logic

import (
	"context"
	"errors"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"
	"log"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// 判断验证码是否一致
	result, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取到该邮箱的验证码")
	}
	if result != req.Code {
		err = errors.New("验证码错误")
		return
	}

	// 判断用户名是否已存在
	count, err := l.svcCtx.SQL.Where("name = ? ", req.Name).Count(new(model.UserInfo))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("用户名已存在")
		return
	}

	// 都无问题，写入数据库
	user := &model.UserInfo{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	n, err := l.svcCtx.SQL.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row: ", n)
	return
}
