package logic

import (
	"context"
	"light-cloud/src/core/define"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"

	"light-cloud/src/core/user/internal/svc"
	"light-cloud/src/core/user/internal/types"

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
	resp = new(types.UserRegisterResponse)

	// 判断用户名和密码
	if len(req.Name) < 6 {
		resp.Msg = "用户名长度不能小于6位"
		return
	}
	if len(req.Password) < 6 {
		resp.Msg = "密码长度不能小于6位"
		return
	}

	// 判断验证码是否一致
	result, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		resp.Msg = "未获取到该邮箱的验证码"
		return
	}
	if result != req.Code {
		resp.Msg = "验证码错误"
		return
	}

	// 判断用户名是否已存在
	count, err := l.svcCtx.SQL.Where("name = ? ", req.Name).Count(new(model.UserInfo))
	if err != nil {
		resp.Msg = "出错了"
		return
	}
	if count > 0 {
		resp.Msg = "用户名已存在"
		return
	}

	// 都无问题，写入数据库
	user := &model.UserInfo{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
		Capacity: define.UserRepositoryMinSize,
	}
	_, err = l.svcCtx.SQL.Insert(user)

	if err != nil {
		resp.Msg = "出错了"
		return
	}

	resp.Msg = "注册成功"
	resp.Code = 200

	return
}
