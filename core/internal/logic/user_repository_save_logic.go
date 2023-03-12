package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
	"light-cloud/src/model"
	"log"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveResponse, err error) {
	ur := &model.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	n, err := l.svcCtx.SQL.Insert(ur)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row: ", n)
	return
}
