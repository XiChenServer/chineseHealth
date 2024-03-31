package logic

import (
	"chineseHealthy/apps/medicine/rpc/types/medicine"
	"context"

	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageDeleteLogic {
	return &ImageDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageDeleteLogic) ImageDelete(req *types.ImageDelRequest) (resp *types.ImageDelResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.MedicineRpc.ImageDelete(l.ctx, &medicine.ImageDelRequest{Id: int64(req.Id)})
	if err != nil {
		return nil, err
	}
	return &types.ImageDelResponse{
		Code:    200,
		Message: "删除成功",
	}, nil

}
