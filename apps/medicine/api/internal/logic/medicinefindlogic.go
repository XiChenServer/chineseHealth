package logic

import (
	"chineseHealthy/apps/medicine/rpc/types/medicine"
	"context"

	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMedicineFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineFindLogic {
	return &MedicineFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MedicineFindLogic) MedicineFind(req *types.MedicineFindRequest) (resp *types.MedicineFindResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.MedicineRpc.MedicineFind(l.ctx, &medicine.MedicineFindRequest{Id: int32(req.Id)})
	if err != nil {
		return nil, err
	}
	medicineInfo := types.MedicineInfo{
		Id:                uint64(res.MedicineInfo.Id),
		Name:              res.MedicineInfo.Name,
		Alias:             res.MedicineInfo.Alias,
		Taste:             res.MedicineInfo.Taste,
		Meridian:          res.MedicineInfo.Meridian,
		Efficacy:          res.MedicineInfo.Efficacy,
		UsageDosage:       res.MedicineInfo.UsageDosage,
		Contraindications: res.MedicineInfo.Contraindications,
	}
	var Info []*types.ImageInfo
	for _, v := range res.MedicineInfo.ImageInfo {
		newInfo := &types.ImageInfo{
			Id:         v.Id,
			MedicineId: v.MedicineId,
			Url:        v.Url,
		}
		Info = append(Info, newInfo)
	}
	medicineInfo.ImageInfo = Info
	return &types.MedicineFindResponse{
		Data: &medicineInfo,
	}, nil
}
