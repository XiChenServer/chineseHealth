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
	var newImage model.MedicineImages
	for _, v := range in.Urls {
		newImage = model.MedicineImages{
			MedicineId: in.MedicineId,
			ImageUrl:   v,
		}

		res, err := l.svcCtx.MedicineImages.Insert(l.ctx, &newImage)
		if err != nil {
			if err == model.ErrNotFound {
				return nil, status.Error(500, err.Error())
			}
		}
		newImage.Id, err = res.LastInsertId()
	}
	return &medicine.ImageCreteResponse{
		Code:    200,
		Message: "图片绑定成功",
		Data: &medicine.ImageInfo{
			Id:         newImage.Id,
			MedicineId: newImage.MedicineId,
			Url:        newImage.ImageUrl,
		},
	}, nil
}
