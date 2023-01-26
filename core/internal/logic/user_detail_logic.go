package logic

import (
	"context"
	"errors"
	"light-cloud/src/core/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	resp = &types.UserDetailResponse{}
	userInfo := new(model.UserInfo)

	get, err := l.svcCtx.SQL.Where("identity = ? ", req.Identity).Get(userInfo)

	if err != nil {
		return nil, err
	}

	if !get {
		return nil, errors.New("user not found")
	}

	// 如果存在
	resp.Name = userInfo.Name
	resp.Email = userInfo.Email

	return
}
