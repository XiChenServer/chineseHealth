package image

import (
	"context"

	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"

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

func (l *ImageUpdataLogic) ImageUpdata() (resp *types.ImageUpdataResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
