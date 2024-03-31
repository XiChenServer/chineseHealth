package logic

import (
	"chineseHealthy/pkg"
	"context"
	"fmt"
	"net/http"

	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageUpdataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageUpdataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageUpdataLogic {
	return &ImageUpdataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageUpdataLogic) ImageUpdata(r *http.Request) (resp *types.ImageUpdataResponse, err error) {
	// todo: add your logic here and delete this line
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	url, err := pkg.UploadToQiNiu(file, handler.Size)
	if err != nil {
		return nil, err
	}

	return &types.ImageUpdataResponse{
		Code:    200,
		Message: "上传成功",
		Data:    url,
	}, nil
}
