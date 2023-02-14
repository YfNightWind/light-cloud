package logic

import (
	"context"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateRequest, userIdentity string) (resp *types.UserUpdateResponse, err error) {
	resp = new(types.UserUpdateResponse)
	userMap := make(map[string]interface{})
	// 用户名更新
	if req.Name != "" {
		exist := l.IsInfoExits("name", req.Name)
		if exist {
			resp.Msg = "用户名已存在"
			return
		}
		userMap["name"] = req.Name
	}

	// 邮箱更新
	if req.Email != "" {
		exist := l.IsInfoExits("email", req.Email)
		if exist {
			resp.Msg = "邮箱已存在"
			return
		}
		userMap["email"] = req.Email
	}

	// 密码更新
	if req.Password != "" {
		userMap["password"] = helper.Md5(req.Password)
	}

	// 头像更新
	if req.Avatar != "" {
		userMap["avatar"] = req.Avatar
	}
	_, err = l.svcCtx.SQL.Table("user_info").Where("identity = ? ", userIdentity).Update(&userMap)
	if err != nil {
		resp.Msg = "出错了"
		return
	}
	resp.Msg = "success"
	return
}

// IsInfoExits 根据字段检查是否已存在
func (l *UserUpdateLogic) IsInfoExits(field string, value string) (exist bool) {
	count, err := l.svcCtx.SQL.Where(field+" = ?", value).Count(new(model.UserInfo))
	if err != nil {
		return
	}
	return count > 0
}
