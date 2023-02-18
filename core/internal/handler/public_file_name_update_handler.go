package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"light-cloud/src/core/internal/logic"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
)

func PublicFileNameUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileNameUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPublicFileNameUpdateLogic(r.Context(), svcCtx)
		resp, err := l.PublicFileNameUpdate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
