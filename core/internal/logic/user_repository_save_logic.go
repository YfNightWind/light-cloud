package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"light-cloud/src/core/define"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
	"light-cloud/src/core/model"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveResponse, err error) {
	resp = new(types.UserRepositorySaveResponse)
	user := &model.UserRepository{
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
		Table("user_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("user_repository.user_identity = ? AND user_repository.deleted_at IS NULL", userIdentity).
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Get(&Size)

	if Size.TotalSize >= define.UserRepositoryMaxSize {
		resp.Msg = "容量不足"
		return
	}

	// 判断是否存在
	count, err := l.svcCtx.SQL.
		Where("name = ? AND parent_id = ? AND user_identity = ? AND deleted_at IS NULL", req.Name, req.ParentId, userIdentity).
		Count(new(model.UserRepository))
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
