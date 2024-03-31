package logic

import (
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

	return
}
