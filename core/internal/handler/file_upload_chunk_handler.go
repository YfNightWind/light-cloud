package handler

import (
	"errors"
	"light-cloud/src/core/helper"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"light-cloud/src/core/internal/logic"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 参数必填判断
		if r.PostForm.Get("key") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("key is empty"))
			return
		}
		if r.PostForm.Get("upload_id") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("upload_id is empty"))
			return
		}
		if r.PostForm.Get("part_number") == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("part_number is empty"))
			return
		}

		etag, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)

		resp = new(types.FileUploadChunkResponse)
		resp.Etag = etag

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
