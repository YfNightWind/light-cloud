package handler

import (
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"light-cloud/src/core/internal/logic"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
)

func FileDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		rp := new(model.RepositoryPool)
		_ = svcCtx.SQL.Where("path = ? ", req.Path).Find(rp)
		// 若文件不存在
		if rp.Id == 0 {
			httpx.OkJson(w, &types.FileDownloadResponse{
				Msg: "file not exits",
			})
			return
		}

		// 一切无误，下载文件
		download, err := helper.FileDownload(r, req.Path, req.Name)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, &types.FileDownloadResponse{
				Data: download,
				Msg:  "success",
			})
		}

		l := logic.NewFileDownloadLogic(r.Context(), svcCtx)
		resp, err := l.FileDownload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
