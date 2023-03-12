package logic

import (
	"context"
	"light-cloud/src/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileDeleteLogic {
	return &PublicFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileDeleteLogic) PublicFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteResponse, err error) {
	_, err = l.svcCtx.SQL.
		Where("user_identity = ? AND identity = ? ", userIdentity, req.Identity).
		Delete(new(model.PublicRepository))
	if err != nil {
		resp.Msg = "删除出错"
		return
	}
	resp = new(types.UserFileDeleteResponse)
	resp.Msg = "删除成功"
	return

	return
}
