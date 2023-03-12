package logic

import (
	"context"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
	"light-cloud/src/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteResponse, err error) {
	resp = new(types.UserFileDeleteResponse)
	_, err = l.svcCtx.SQL.
		Where("user_identity = ? AND identity = ? ", userIdentity, req.Identity).
		Delete(new(model.UserRepository))
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
