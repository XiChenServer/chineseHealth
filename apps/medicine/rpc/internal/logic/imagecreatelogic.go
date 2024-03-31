package logic

import (
	"chineseHealthy/apps/medicine/model"
	"context"
	"google.golang.org/grpc/status"

	"chineseHealthy/apps/medicine/rpc/internal/svc"
	"chineseHealthy/apps/medicine/rpc/types/medicine"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImageCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageCreateLogic {
	return &ImageCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ImageCreateLogic) ImageCreate(in *medicine.ImageCreateRequest) (*medicine.ImageCreteResponse, error) {
	// todo: add your logic here and delete this line
	for _, v := range in.Urls {
		newImage := model.MedicineImages{
			MedicineId: in.MedicineId,
			ImageUrl:   v,
		}
		_, err := l.svcCtx.MedicineImages.Insert(l.ctx, &newImage)
		if err != nil {
			if err == model.ErrNotFound {
				return nil, status.Error(500, err.Error())
			}
		}
	}
	return &medicine.ImageCreteResponse{
		Code:    0,
		Message: "",
		Data:    nil,
	}, nil
}
