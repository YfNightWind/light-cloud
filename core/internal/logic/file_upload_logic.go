package logic

import (
	"context"
	"light-cloud/src/core/helper"
	"light-cloud/src/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	resp = new(types.FileUploadResponse)
	rp := &model.RepositoryPool{
		Identity: helper.UUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}

	_, err = l.svcCtx.SQL.
		Select("identity, name, hash, path, ext, size, created_at, updated_at").
		Insert(rp)
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Identity = rp.Identity
	resp.Name = rp.Name
	resp.Ext = rp.Ext
	resp.Msg = "success"
	return
}
