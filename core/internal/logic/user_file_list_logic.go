package logic

import (
	"context"
	"light-cloud/src/core/define"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
	"light-cloud/src/core/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	resp = new(types.UserFileListResponse)
	userFile := make([]*types.UserFile, 0)
	deletedFile := make([]*types.DeletedUserFile, 0)

	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}

	page := req.Page
	if page == 0 {
		page = 1
	}

	//offset := (page - 1) * size

	// 查询用户文件列表
	//l.svcCtx.SQL.ShowSQL(true)
	err = l.svcCtx.SQL.
		Table("user_repository").
		Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.updated_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("user_identity = ?", userIdentity).
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		//Limit(size, offset).
		Find(&userFile)
	if err != nil {
		resp.Msg = "error"
		return
	}

	// 查询用户已删除文件
	err = l.svcCtx.SQL.
		Table("user_repository").
		Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.deleted_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("user_identity = ? ", userIdentity).
		Where("user_repository.deleted_at IS NOT NULL").
		// Order("user_repository.deleted_at desc").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Find(&deletedFile)

	if err != nil {
		resp.Msg = "error"
		return
	}

	// 查询总数
	count, err := l.svcCtx.SQL.
		Where("parent_id = ? AND user_identity = ? AND deleted_at IS NULL", req.Id, userIdentity).
		Count(new(model.UserRepository))
	if err != nil {
		return
	}

	resp.List = userFile
	resp.DeletedList = deletedFile
	resp.Count = count
	resp.Msg = "success"
	return
}
