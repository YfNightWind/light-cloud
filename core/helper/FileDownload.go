package helper

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"light-cloud/src/core/define"
	"net/http"
	"net/url"
)

// FileDownload 从腾讯云下载(高级接口）
// 当对象大于16MB时，采用 Range 方式下载文件
func FileDownload(r *http.Request, path string, fileName string) ([]byte, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	opt := &cos.MultiDownloadOptions{
		ThreadPoolSize: 5,
	}

	response, err := client.Object.Download(context.Background(), fileName, path, opt)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, err
}
