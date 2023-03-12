package logic

import (
	"context"
	"light-cloud/src/core/define"
	"light-cloud/src/core/helper"
	"light-cloud/src/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileSaveLogic {
	return &PublicFileSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileSaveLogic) PublicFileSave(req *types.PublicRepositorySaveRequest, userIdentity string) (resp *types.PublicRepositorySaveResponse, err error) {
	resp = new(types.PublicRepositorySaveResponse)
	user := model.PublicRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}

	// 判断容量
	var Size struct {
		TotalSize int `json:"total_size"`
	}
	_, err = l.svcCtx.SQL.
		Table("public_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("public_repository.user_identity = ? AND public_repository.deleted_at IS NULL", userIdentity).
		Join("LEFT", "repository_pool", "public_repository.repository_identity = repository_pool.identity").
		Get(&Size)
	if err != nil {
		resp.Msg = "error"
		return
	}
	if userIdentity != "USER_1" && Size.TotalSize >= define.PublicRepositoryMaxSize {
		resp.Msg = "容量不足"
		return
	}

	// 判断是否存在
	count, err := l.svcCtx.SQL.
		Where("name = ? AND parent_id = ? AND user_identity = ? AND deleted_at IS NULL", req.Name, req.ParentId, userIdentity).
		Count(new(model.PublicRepository))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		resp.Msg = "exist"
		resp.Code = 405
		return
	}

	// 一切正常，保存
	_, err = l.svcCtx.SQL.Insert(user)
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Msg = "success"
	resp.Code = 200
	return
}
