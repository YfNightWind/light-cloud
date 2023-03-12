package logic

import (
	"context"
	"light-cloud/src/core/define"
	"light-cloud/src/model"
	"time"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileListLogic {
	return &PublicFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileListLogic) PublicFileList(req *types.PublicFileListRequest) (resp *types.PublicFileListResponse, err error) {
	resp = new(types.PublicFileListResponse)
	publicFile := make([]*types.PublicFile, 0)

	//// 分页参数
	//size := req.Size
	//if size == 0 {
	//	size = define.PageSize
	//}
	//
	//page := req.Page
	//if page == 0 {
	//	page = 1
	//}

	err = l.svcCtx.SQL.
		Table("public_repository").
		Select("public_repository.id, public_repository.parent_id, public_repository.identity, "+
			"public_repository.repository_identity, public_repository.ext, public_repository.updated_at,"+
			"public_repository.name, repository_pool.path, repository_pool.size, user_info.name as owner").
		Where("public_repository.deleted_at = ? OR public_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Join("LEFT", "repository_pool", "public_repository.repository_identity = repository_pool.identity").
		Join("LEFT", "user_info", "public_repository.user_identity = user_info.identity").
		Find(&publicFile)
	if err != nil {
		resp.Msg = "error"
		return
	}

	// 查询总数
	count, err := l.svcCtx.SQL.
		Table("public_repository").
		Where("deleted_at IS NULL").
		Count(new(model.PublicRepository))
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.Msg = "success"
	resp.Count = count
	resp.List = publicFile

	return
}
