package logic

import (
	"context"
	"light-cloud/src/core/define"
	"light-cloud/src/core/helper"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationResponse, err error) {
	resp = new(types.RefreshAuthorizationResponse)
	userClaim, err := helper.AnalyzeToken(authorization)
	if err != nil {
		resp.Msg = "解析token失败"
		return
	}

	token, err := helper.GenerateToken(userClaim.Id, userClaim.Identity, userClaim.Name, define.TokenExpire)
	if err != nil {
		resp.Msg = "生成token失败"
		return
	}

	refreshToken, err := helper.GenerateToken(userClaim.Id, userClaim.Identity, userClaim.Name, define.RefreshTokenExpire)
	if err != nil {
		resp.Msg = "生成token失败"
		return
	}

	resp.Msg = "success"
	resp.Token = token
	resp.RefreshToken = refreshToken

	return
}
