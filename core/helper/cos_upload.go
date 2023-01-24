package helper

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"light-cloud/src/core/define"
	"net/http"
	"path"
)

// CosUpload 上传文件到七牛云
func CosUpload(r *http.Request, fileSize int64) (string, error) {
	bucket := define.Bucket

	// 获取上传文件的对象
	file, fileHeader, err := r.FormFile("file")

	key := "light-cloud/" + UUID() + path.Ext(fileHeader.Filename)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(define.AccessKey, define.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}

	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong

	// 是否使用https域名
	cfg.UseHTTPS = true

	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	err = formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, nil)
	if err != nil {
		return "", err
	}

	return key, nil
}
