package logic

import (
	"context"
	"light-cloud/src/core/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateResponse, err error) {
	resp = new(types.UserFileNameUpdateResponse)

	if req.Name == "" {
		resp.Msg = "文件名为空"
		return
	}

	// 判断该名称在当前文件夹下是否存在
	count, err := l.svcCtx.SQL.
		Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository as ur WHERE ur.identity = ? )", req.Name, req.Identity).
		Count(new(model.UserRepository))
	if err != nil {
		resp.Msg = "error"
		return
	}
	if count > 0 {
		resp.Msg = "文件名已存在"
		return
	}

	// 文件名称修改
	data := model.UserRepository{Name: req.Name}
	_, err = l.svcCtx.SQL.Where("identity = ? AND user_identity = ? ", req.Identity, userIdentity).Update(data)
	if err != nil {
		resp.Msg = "error"
		return
	}
	resp.Msg = "success"
	return
}
