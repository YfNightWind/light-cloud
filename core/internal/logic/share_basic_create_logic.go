package logic

import (
	"context"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"

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
	l.svcCtx.SQL.ShowSQL(true)
	data := model.ShareBasic{
		Identity:           uuid,
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		ExpiredTime:        req.ExpiredTime,
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
