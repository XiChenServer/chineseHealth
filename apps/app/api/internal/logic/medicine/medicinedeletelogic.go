package medicine

import (
	"context"

	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMedicineDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineDeleteLogic {
	return &MedicineDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MedicineDeleteLogic) MedicineDelete(req *types.MedicineDelRequest) (resp *types.MedicineDelResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
