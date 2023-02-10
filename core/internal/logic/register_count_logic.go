package logic

import (
	"context"
	"light-cloud/src/core/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCountLogic {
	return &RegisterCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterCountLogic) RegisterCount(req *types.RegisterCountRequest) (resp *types.RegisterCountResponse, err error) {
	resp = new(types.RegisterCountResponse)
	count, err := l.svcCtx.SQL.Where("deleted_at IS NULL").Count(new(model.UserInfo))
	if err != nil {
		resp.Msg = "出错了"
		return
	}
	resp.Msg = "success"
	resp.Count = count

	return
}
