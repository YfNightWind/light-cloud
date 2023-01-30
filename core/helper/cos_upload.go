package helper

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"light-cloud/src/core/define"
	"net/http"
	"net/url"
	"path"
)

// CosUpload 上传文件到腾讯云
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	fileName := UUID() + path.Ext(fileHeader.Filename)
	key := "/" + fileName

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}

	// 返回文件链接
	return define.CosBucket + key, nil
}
