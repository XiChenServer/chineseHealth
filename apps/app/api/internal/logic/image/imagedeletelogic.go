package image

import (
	"context"

	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageDeleteLogic {
	return &ImageDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageDeleteLogic) ImageDelete(req *types.ImageDelRequest) (resp *types.ImageDelResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
