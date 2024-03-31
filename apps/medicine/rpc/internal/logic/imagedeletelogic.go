package logic

import (
	"context"

	"chineseHealthy/apps/medicine/rpc/internal/svc"
	"chineseHealthy/apps/medicine/rpc/types/medicine"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageDeleteLogic {
	return &ImageDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ImageDeleteLogic) ImageDelete(in *medicine.ImageDelRequest) (*medicine.ImageDelResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.MedicineImages.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &medicine.ImageDelResponse{
		Code:    200,
		Message: "删除成功",
	}, nil
}
