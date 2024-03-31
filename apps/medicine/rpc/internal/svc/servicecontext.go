package svc

import (
	"chineseHealthy/apps/medicine/model"
	"chineseHealthy/apps/medicine/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	MedicineInfo   model.ChineseMedicineInfoModel
	MedicineImages model.MedicineImagesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		MedicineInfo:   model.NewChineseMedicineInfoModel(conn, c.CacheRedis),
		MedicineImages: model.NewMedicineImagesModel(conn, c.CacheRedis),
	}
}
