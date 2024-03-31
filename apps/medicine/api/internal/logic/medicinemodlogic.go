package logic

import (
	"context"

	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"

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

	return
}
