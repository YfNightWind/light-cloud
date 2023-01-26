package logic

import (
	"context"
	"light-cloud/src/core/define"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
	"light-cloud/src/core/model"

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

	offset := (page - 1) * size

	// 查询用户文件列表
	err = l.svcCtx.SQL.
		Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext, "+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Limit(size, offset).
		Find(&userFile)
	if err != nil {
		return
	}

	// 查询总数
	count, err := l.svcCtx.SQL.
		Where("parent_id = ? AND user_identity = ? ", req.Id, userIdentity).Count(new(model.UserRepository))
	if err != nil {
		return
	}

	resp.List = userFile
	resp.DeletedList = deletedFile
	resp.Count = count
	return
}
