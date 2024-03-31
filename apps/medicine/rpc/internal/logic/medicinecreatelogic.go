package logic

import (
	"chineseHealthy/apps/medicine/model"
	"context"
	"database/sql"
	"google.golang.org/grpc/status"

	"chineseHealthy/apps/medicine/rpc/internal/svc"
	"chineseHealthy/apps/medicine/rpc/types/medicine"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMedicineCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineCreateLogic {
	return &MedicineCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MedicineCreateLogic) MedicineCreate(in *medicine.MedicineCreateRequest) (*medicine.MedicineCreateResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.MedicineInfo.FindOneByName(l.ctx, in.MedicineInfo.Name)
	if err == nil {
		return nil, status.Error(100, "该药品已存在")
	}
	var newMedicine model.ChineseMedicineInfo

	var medicineInfo medicine.ChineseMedicineInfo
	if err == model.ErrNotFound {
		newMedicine = model.ChineseMedicineInfo{
			Name:              in.MedicineInfo.Name,
			Alias:             sql.NullString{String: in.MedicineInfo.Alias, Valid: true},
			Taste:             sql.NullString{String: in.MedicineInfo.Taste, Valid: true},
			Meridian:          sql.NullString{String: in.MedicineInfo.Meridian, Valid: true},
			Efficacy:          sql.NullString{String: in.MedicineInfo.Efficacy, Valid: true},
			UsageDosage:       sql.NullString{String: in.MedicineInfo.UsageDosage, Valid: true},
			Contraindications: sql.NullString{String: in.MedicineInfo.Contraindications, Valid: true},
		}
		res, err := l.svcCtx.MedicineInfo.Insert(l.ctx, &newMedicine)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newMedicine.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

	}

	medicineInfo = medicine.ChineseMedicineInfo{
		Id:                int32(newMedicine.Id),
		Name:              newMedicine.Name,
		Alias:             in.MedicineInfo.Alias,
		Taste:             in.MedicineInfo.Taste,
		Meridian:          in.MedicineInfo.Meridian,
		Efficacy:          in.MedicineInfo.Efficacy,
		UsageDosage:       in.MedicineInfo.UsageDosage,
		Contraindications: in.MedicineInfo.Contraindications,
	}

	return &medicine.MedicineCreateResponse{
		Code:         200,
		Message:      "创建成功",
		MedicineInfo: &medicineInfo,
	}, nil
}
