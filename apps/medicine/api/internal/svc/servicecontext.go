package svc

import (
	"chineseHealthy/apps/medicine/api/internal/config"
	"chineseHealthy/apps/medicine/rpc/medicineclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	MedicineRpc medicineclient.Medicine
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MedicineRpc: medicineclient.NewMedicine(zrpc.MustNewClient(c.MedicineRpc)),
	}
}
