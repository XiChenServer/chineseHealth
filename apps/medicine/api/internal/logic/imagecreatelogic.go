package logic

import (
	"context"

	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageCreateLogic {
	return &ImageCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageCreateLogic) ImageCreate(req *types.ImageCreateRequest) (resp *types.ImageCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
