package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"chineseHealthy/apps/medicine/rpc/internal/svc"
	"chineseHealthy/apps/medicine/rpc/types/medicine"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineModLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMedicineModLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineModLogic {
	return &MedicineModLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MedicineModLogic) MedicineMod(in *medicine.MedicineModRequest) (*medicine.MedicineModResponse, error) {
	// todo: add your logic here and delete this line
	// 检查药品是否已存在
	existingMedicine, err := l.svcCtx.MedicineInfo.FindOneByName(l.ctx, in.MedicineInfo.Name)
	if err != nil {
		return nil, status.Error(100, "该药品不存在")
	}
	if in.MedicineInfo.Name != "" {
		existingMedicine.Name = in.MedicineInfo.Name
	}
	if in.MedicineInfo.Alias != "" {
		existingMedicine.Alias.String = in.MedicineInfo.Alias
	}
	if in.MedicineInfo.Meridian != "" {
		existingMedicine.Meridian.String = in.MedicineInfo.Meridian
	}
	if in.MedicineInfo.Efficacy != "" {
		existingMedicine.Efficacy.String = in.MedicineInfo.Efficacy
	}
	if in.MedicineInfo.UsageDosage != "" {
		existingMedicine.UsageDosage.String = in.MedicineInfo.UsageDosage
	}
	if in.MedicineInfo.Contraindications != "" {
		existingMedicine.UsageDosage.String = in.MedicineInfo.UsageDosage
	}
	image, err := l.svcCtx.MedicineImages.FindAllImage(l.ctx, int64(in.Id))
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	var medicineInfo = medicine.ChineseMedicineInfo{
		Id:                int32(existingMedicine.Id),
		Name:              existingMedicine.Name,
		Alias:             existingMedicine.Alias.String,
		Taste:             existingMedicine.Taste.String,
		Meridian:          existingMedicine.Meridian.String,
		Efficacy:          existingMedicine.Efficacy.String,
		UsageDosage:       existingMedicine.UsageDosage.String,
		Contraindications: existingMedicine.Contraindications.String,
		ImageUrls:         image,
	}
	return &medicine.MedicineModResponse{
		Code:         200,
		Message:      "修改成功",
		MedicineInfo: &medicineInfo,
	}, nil
}
