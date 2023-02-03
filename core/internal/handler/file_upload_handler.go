package handler

import (
	"crypto/md5"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/internal/logic"
	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"
	"light-cloud/src/core/model"
	"net/http"
	"path"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}

		// 判断文件是否已存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b)) // 将返回的十六进制转为字符串作为文件的Hash值
		rp := new(model.RepositoryPool)
		get, err := svcCtx.SQL.Where("hash = ? ", hash).Get(rp)
		if err != nil {
			return
		}
		if get {
			httpx.OkJson(w, &types.FileUploadResponse{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}

		// 如果不存在，往腾讯云中存储
		url, err := helper.CosUpload(r)
		if err != nil {
			return
		}

		// 往 logic 中传递 request
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = url

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
