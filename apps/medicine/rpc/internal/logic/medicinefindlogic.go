package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"chineseHealthy/apps/medicine/rpc/internal/svc"
	"chineseHealthy/apps/medicine/rpc/types/medicine"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMedicineFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineFindLogic {
	return &MedicineFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MedicineFindLogic) MedicineFind(in *medicine.MedicineFindRequest) (*medicine.MedicineFindResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.MedicineInfo.FindOne(l.ctx, int64(in.Id))
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	image, err := l.svcCtx.MedicineImages.FindAllImage(l.ctx, int64(in.Id))
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	var medicineInfo = medicine.ChineseMedicineInfo{
		Id:                int32(res.Id),
		Name:              res.Name,
		Alias:             res.Alias.String,
		Taste:             res.Taste.String,
		Meridian:          res.Meridian.String,
		Efficacy:          res.Efficacy.String,
		UsageDosage:       res.UsageDosage.String,
		Contraindications: res.Contraindications.String,
		ImageUrls:         image,
	}

	return &medicine.MedicineFindResponse{
		Code:         200,
		Message:      "查找信息如下",
		MedicineInfo: &medicineInfo,
	}, nil
}
