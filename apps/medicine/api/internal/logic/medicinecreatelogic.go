package logic

import (
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

	return
}
