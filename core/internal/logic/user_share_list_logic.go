package logic

import (
	"context"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserShareListLogic {
	return &UserShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserShareListLogic) UserShareList(req *types.UserShareListRequest, userIdentity string) (resp *types.UserShareListResponse, err error) {
	shareFile := make([]*types.ShareBasicDetailResponse, 0)
	resp = new(types.UserShareListResponse)

	err = l.svcCtx.SQL.Table("share_basic").
		Select("share_basic.identity, share_basic.repository_identity, user_repository.name, repository_pool.ext, "+
			"repository_pool.path, repository_pool.size, share_basic.click_num, share_basic.desc, "+
			"user_info.name as owner, user_info.avatar, share_basic.expired_time, share_basic.updated_at").
		Join("LEFT", "repository_pool", "repository_pool.identity = share_basic.repository_identity").
		Join("LEFT", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Join("LEFT", "user_info", "share_basic.user_identity = user_info.identity").
		Where("share_basic.user_identity = ? ", userIdentity).
		Where("share_basic.deleted_at IS NULL").
		OrderBy("share_basic.updated_at desc").
		Find(&shareFile)

	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.List = shareFile
	resp.Msg = "success"
	resp.Code = 0
	return
}
