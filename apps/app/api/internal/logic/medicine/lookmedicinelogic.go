package medicine

import (
	"chineseHealthy/apps/medicine/rpc/types/medicine"
	"context"

	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookMedicineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookMedicineLogic {
	return &LookMedicineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookMedicineLogic) LookMedicine(req *types.LookMedicineVagueRequest) (resp *types.LookMedicineVagueResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.MedicineRpc.FindMedicine(l.ctx, &medicine.FindMedicineRequest{Name: req.Key})
	if err != nil {
		return nil, err
	}
	var medicine []*types.MedicineInfo

	for _, v := range res.MedicineInfo {
		newMedicine := &types.MedicineInfo{
			Id:                uint64(v.Id),
			Name:              v.Name,
			Alias:             v.Alias,
			Taste:             v.Taste,
			Meridian:          v.Meridian,
			Efficacy:          v.Efficacy,
			UsageDosage:       v.UsageDosage,
			Contraindications: v.Contraindications,
		}
		var Info []*types.ImageInfo
		for _, i := range v.ImageInfo {
			newInfo := &types.ImageInfo{
				Id:         i.Id,
				MedicineId: i.MedicineId,
				Url:        i.Url,
			}
			Info = append(Info, newInfo)
		}
		newMedicine.ImageInfo = Info
		medicine = append(medicine, newMedicine)

	}

	return &types.LookMedicineVagueResponse{Data: medicine}, nil
}
