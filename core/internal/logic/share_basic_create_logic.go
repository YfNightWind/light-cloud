package logic

import (
	"context"
	"light-cloud/src/core/helper"
	model2 "light-cloud/src/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	resp = new(types.ShareBasicCreateResponse)
	uuid := helper.UUID()

	// 获取用户池子中文件的identity
	usr := new(model2.UserRepository)
	get, err := l.svcCtx.SQL.Table("user_repository").
		Where("identity = ? ", req.UserRepositoryIdentity).
		Get(usr)
	if !get {
		resp.Msg = "user repository resource not found"
		return
	}
	if err != nil {
		resp.Msg = "error"
		return
	}

	data := model2.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     usr.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	_, err = l.svcCtx.SQL.Insert(data)
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Identity = uuid
	resp.Msg = "success"
	return
}
