package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChineseMedicineInfoModel = (*customChineseMedicineInfoModel)(nil)

type (
	// ChineseMedicineInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChineseMedicineInfoModel.
	ChineseMedicineInfoModel interface {
		chineseMedicineInfoModel
	}

	customChineseMedicineInfoModel struct {
		*defaultChineseMedicineInfoModel
	}
)

// NewChineseMedicineInfoModel returns a model for the database table.
func NewChineseMedicineInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChineseMedicineInfoModel {
	return &customChineseMedicineInfoModel{
		defaultChineseMedicineInfoModel: newChineseMedicineInfoModel(conn, c, opts...),
	}
}
