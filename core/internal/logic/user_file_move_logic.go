package logic

import (
	"context"
	"light-cloud/src/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	resp = new(types.UserFileMoveResponse)

	// parent identity
	parentData := new(model.UserRepository)
	get, err := l.svcCtx.SQL.Table("user_repository").
		Where("identity = ? AND user_identity = ? ", req.ParentIdentity, userIdentity).
		Get(parentData)
	if !get {
		resp.Msg = "文件夹不存在"
		return
	}
	if err != nil {
		resp.Msg = "error"
		return
	}

	// 更新记录 parent identity
	_, err = l.svcCtx.SQL.Table("user_repository").
		Where("identity = ? AND deleted_at IS NULL", req.Identity).
		Update(model.UserRepository{
			ParentId: int64(parentData.Id),
		})
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
