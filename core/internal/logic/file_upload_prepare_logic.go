package logic

import (
	"context"
	"light-cloud/src/core/helper"
	"light-cloud/src/core/model"

	"light-cloud/src/core/internal/svc"
	"light-cloud/src/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareResponse, err error) {
	rp := new(model.RepositoryPool)
	resp = new(types.FileUploadPrepareResponse)

	get, err := l.svcCtx.SQL.Where("hash = ? ", req.Md5).Get(rp)
	if err != nil {
		resp.Msg = "error"
		return
	}
	if get {
		// 文件存在，秒传成功
		resp.Identity = rp.Identity
		resp.Msg = "秒传成功"
	} else {
		// 获取文件的 UploadID 和 key 进行COS文件分片上传
		key, uploadID, err := helper.CosInitPart(req.Ext)
		if err != nil {
			resp.Msg = "error"
			return resp, err
		}
		resp.UploadId = uploadID
		resp.Key = key
		resp.Code = 200
		resp.Msg = "success"
	}

	return
}
