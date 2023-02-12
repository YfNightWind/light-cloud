package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"light-cloud/src/core/internal/logic"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
)

func ShareStatisticsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareStatisticsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShareStatisticsLogic(r.Context(), svcCtx)
		resp, err := l.ShareStatistics(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
