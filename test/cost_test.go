package test

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"light-cloud/src/core/define"
	"testing"
)

func TestFileUploadByFilepath(t *testing.T) {
	localFile := "你的本地地址/xxx.png"
	bucket := define.Bucket
	key := "light-cloud/test.png"
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

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
