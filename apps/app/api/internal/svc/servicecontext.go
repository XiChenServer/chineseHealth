package svc

import (
	"chineseHealthy/apps/app/api/internal/config"
	"chineseHealthy/apps/article/rpc/articleclient"
	"chineseHealthy/apps/medicine/rpc/medicineclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	MedicineRpc medicineclient.Medicine
	ArticleRpc  articleclient.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MedicineRpc: medicineclient.NewMedicine(zrpc.MustNewClient(c.MedicineRpc)),
		ArticleRpc:  articleclient.NewArticle(zrpc.MustNewClient(c.ArticleRpc)),
	}
}
