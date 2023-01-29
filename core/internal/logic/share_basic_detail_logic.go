package logic

import (
	"context"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailResponse, err error) {
	resp = new(types.ShareBasicDetailResponse)
	// 更新分享记录的点击次数，每次分享点击次数 + 1
	_, err = l.svcCtx.SQL.
		Table("share_basic").
		Exec("UPDATE share_basic SET click_num = click_num + 1 where identity = ?", req.Identity)
	if err != nil {
		resp.Msg = "error"
		return
	}

	// 获取资源详细信息
	_, err = l.svcCtx.SQL.
		Table("share_basic").
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Join("LEFT", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ? ", req.Identity).
		Get(resp)
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"

	return
}
