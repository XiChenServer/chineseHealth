package logic

import (
	"chineseHealthy/apps/medicine/rpc/types/medicine"
	"context"

	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"

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
	_, err = l.svcCtx.MedicineRpc.MedicineDel(l.ctx, &medicine.MedicineDelRequest{Id: int32(req.Id)})
	if err != nil {
		return nil, err
	}
	return &types.MedicineDelResponse{
		Code:    200,
		Message: "删除成功",
	}, nil
}
