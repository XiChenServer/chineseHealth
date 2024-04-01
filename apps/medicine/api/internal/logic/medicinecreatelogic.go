package logic

import (
	"chineseHealthy/apps/medicine/rpc/types/medicine"
	"context"

	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMedicineCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineCreateLogic {
	return &MedicineCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MedicineCreateLogic) MedicineCreate(req *types.MedicineCreateRequest) (resp *types.MedicineCreateResponse, err error) {
	// todo: add your logic here and delete this line
	medicineInfo := medicine.ChineseMedicineInfo{
		Id:                0,
		Name:              req.Data.Name,
		Alias:             req.Data.Alias,
		Taste:             req.Data.Taste,
		Meridian:          req.Data.Meridian,
		Efficacy:          req.Data.Efficacy,
		UsageDosage:       req.Data.UsageDosage,
		Contraindications: req.Data.Contraindications,
	}
	res, err := l.svcCtx.MedicineRpc.MedicineCreate(l.ctx, &medicine.MedicineCreateRequest{MedicineInfo: &medicineInfo})
	if err != nil {
		return nil, err
	}
	return &types.MedicineCreateResponse{
		Code:    200,
		Message: "创建成功",
		Data: types.MedicineInfo{
			Id:                uint64(res.MedicineInfo.Id),
			Name:              res.MedicineInfo.Name,
			Alias:             res.MedicineInfo.Alias,
			Taste:             res.MedicineInfo.Taste,
			Meridian:          res.MedicineInfo.Meridian,
			Efficacy:          res.MedicineInfo.Efficacy,
			UsageDosage:       res.MedicineInfo.UsageDosage,
			Contraindications: res.MedicineInfo.Contraindications,
			ImageUrls:         nil,
		},
	}, nil
}
