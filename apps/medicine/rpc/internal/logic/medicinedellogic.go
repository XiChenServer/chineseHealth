package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"chineseHealthy/apps/medicine/rpc/internal/svc"
	"chineseHealthy/apps/medicine/rpc/types/medicine"

	"github.com/zeromicro/go-zero/core/logx"
)

type MedicineDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMedicineDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MedicineDelLogic {
	return &MedicineDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MedicineDelLogic) MedicineDel(in *medicine.MedicineDelRequest) (*medicine.MedicineDelResponse, error) {
	err := l.svcCtx.MedicineInfo.Delete(l.ctx, int64(in.Id))
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &medicine.MedicineDelResponse{
		Code:    200,
		Message: "删除成功",
	}, nil
}
