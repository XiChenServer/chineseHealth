package logic

import (
	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"
	"chineseHealthy/apps/medicine/rpc/types/medicine"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineModLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMedicineModLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineModLogic {
	return &MedicineModLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MedicineModLogic) MedicineMod(req *types.MedicineModRequest) (resp *types.MedicineModResponse, err error) {
	// todo: add your logic here and delete this line

	medicineInfo := medicine.ChineseMedicineInfo{
		Id:                int32(req.Data.Id),
		Name:              req.Data.Name,
		Alias:             req.Data.Alias,
		Taste:             req.Data.Taste,
		Meridian:          req.Data.Meridian,
		Efficacy:          req.Data.Efficacy,
		UsageDosage:       req.Data.UsageDosage,
		Contraindications: req.Data.Contraindications,
	}
	res, err := l.svcCtx.MedicineRpc.MedicineMod(l.ctx, &medicine.MedicineModRequest{MedicineInfo: &medicineInfo})
	if err != nil {
		return nil, err
	}

	return &types.MedicineModResponse{
		Code:    200,
		Message: "修改成功",
		Data: &types.MedicineInfo{
			Id:                uint64(res.MedicineInfo.Id),
			Name:              res.MedicineInfo.Name,
			Alias:             res.MedicineInfo.Alias,
			Taste:             res.MedicineInfo.Taste,
			Meridian:          res.MedicineInfo.Meridian,
			Efficacy:          res.MedicineInfo.Efficacy,
			UsageDosage:       res.MedicineInfo.UsageDosage,
			Contraindications: res.MedicineInfo.Contraindications,
			ImageUrls:         res.MedicineInfo.ImageUrls,
		},
	}, nil
}
