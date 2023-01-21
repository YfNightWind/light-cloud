package logic

import (
	"context"
	"errors"
	"light-cloud/src/core/define"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"
	"time"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendResponse, err error) {
	// 若该邮箱未注册
	count, err := l.svcCtx.Engine.Where("email = ? ", req.Email).Table(model.UserInfo{}).Count()
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}

	// if count == 0⬇️
	// 获取验证码
	code := helper.GenValidateCode()
	// 存储验证码Redis
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.ExpireTime))

	// 发送验证码
	err = helper.SendMailCode(req.Email, code)
	if err != nil {
		return nil, err
	}

	return
}
