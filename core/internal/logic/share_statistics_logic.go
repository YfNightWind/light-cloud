package logic

import (
	"context"
	"light-cloud/src/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareStatisticsLogic {
	return &ShareStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareStatisticsLogic) ShareStatistics(req *types.ShareStatisticsRequest) (resp *types.ShareStatisticsResponse, err error) {
	resp = new(types.ShareStatisticsResponse)

	count, err := l.svcCtx.SQL.Where("deleted_at IS NULL").Count(new(model.ShareBasic))
	if err != nil {
		resp.Msg = "error"
		return
	}

	clickNum, err := l.svcCtx.SQL.
		Table("share_basic").
		Where("deleted_at IS NULL").
		SumInt(types.ShareStatisticsResponse{}, "click_num")
	if err != nil {
		resp.Msg = "error"
		return
	}

	resp.ShareCount = int(count)
	resp.ClickNum = int(clickNum)
	resp.Msg = "success"
	resp.Code = 200
	return
}
