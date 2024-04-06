package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"chineseHealthy/apps/medicine/rpc/internal/svc"
	"chineseHealthy/apps/medicine/rpc/types/medicine"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMedicineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMedicineLogic {
	return &FindMedicineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMedicineLogic) FindMedicine(in *medicine.FindMedicineRequest) (*medicine.FindMedicineResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.MedicineInfo.FindMedicineVague(l.ctx, in.Name)
	if err != nil {
		return nil, err
	}
	var medicineInfo []*medicine.ChineseMedicineInfo
	for _, v := range *res {
		image, err := l.svcCtx.MedicineImages.FindAllImage(l.ctx, v.Id)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newMedicineInfo := medicine.ChineseMedicineInfo{
			Id:                int32(v.Id),
			Name:              v.Name,
			Alias:             v.Alias.String,
			Taste:             v.Taste.String,
			Meridian:          v.Meridian.String,
			Efficacy:          v.Efficacy.String,
			UsageDosage:       v.UsageDosage.String,
			Contraindications: v.Contraindications.String,
			ImageInfo:         image,
		}
		medicineInfo = append(medicineInfo, &newMedicineInfo)
	}
	return &medicine.FindMedicineResponse{MedicineInfo: medicineInfo}, nil
}
