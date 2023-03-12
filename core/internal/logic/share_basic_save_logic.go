package logic

import (
	"context"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
	model2 "light-cloud/src/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveResponse, err error) {
	resp = new(types.ShareBasicSaveResponse)
	// 获取资源详情 from repository_pool
	rp := new(model2.RepositoryPool)
	get, err := l.svcCtx.SQL.
		Table("repository_pool").
		Where("identity = ?", req.RepositoryIdentity).
		Get(rp)
	if err != nil {
		resp.Msg = "error"
		return
	}
	if !get {
		resp.Msg = "资源不存在"
		return
	}

	// TODO 容量不足判断

	// user_repository 资源保存
	usr := &model2.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	_, err = l.svcCtx.SQL.Insert(usr)
	if err != nil {
		resp.Msg = "save error"
		return
	}
	resp.Identity = usr.Identity
	resp.Msg = "success"
	return
}
