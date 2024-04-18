package medicine

import (
	"chineseHealthy/apps/medicine/rpc/types/medicine"
	"context"
	"fmt"

	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllMedicineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookAllMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllMedicineLogic {
	return &LookAllMedicineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookAllMedicineLogic) LookAllMedicine(req *types.ViewAllResquest) (resp *types.ViewAllResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.MedicineRpc.ViewAll(l.ctx, &medicine.ViewAllRequest{})
	if err != nil {
		return nil, err
	}
	var medicine []*types.MedicineInfo
	fmt.Println("123")
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
		fmt.Println(medicine)
	}
	return &types.ViewAllResponse{
		Data: medicine,
	}, nil
}
