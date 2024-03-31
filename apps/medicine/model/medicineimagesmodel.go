package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MedicineImagesModel = (*customMedicineImagesModel)(nil)

type (
	// MedicineImagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMedicineImagesModel.
	MedicineImagesModel interface {
		medicineImagesModel
	}

	customMedicineImagesModel struct {
		*defaultMedicineImagesModel
	}
)

// NewMedicineImagesModel returns a model for the database table.
func NewMedicineImagesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MedicineImagesModel {
	return &customMedicineImagesModel{
		defaultMedicineImagesModel: newMedicineImagesModel(conn, c, opts...),
	}
}
