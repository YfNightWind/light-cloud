package logic

import (
	"context"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFolderCreateLogic {
	return &PublicFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFolderCreateLogic) PublicFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	resp = new(types.UserFolderCreateResponse)
	if req.Name == "" {
		resp.Msg = "name is empty"
		return
	}
	// 判断该名称在当前文件夹下是否存在
	count, err := l.svcCtx.SQL.
		Where("name = ? AND parent_id = ? AND user_identity = ? AND deleted_at IS NULL", req.Name, req.ParentId, userIdentity).
		Count(new(model.PublicRepository))
	if err != nil {
		resp.Msg = "error"
		return
	}
	if count > 0 {
		resp.Msg = "exits"
		return
	}

	// 创建文件夹
	data := &model.PublicRepository{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.SQL.Insert(data)
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return

	return
}
