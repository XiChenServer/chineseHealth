package svc

import (
	"chineseHealthy/apps/article/rpc/internal/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open("mysql", getDSN(&c))

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

func getDSN(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.Database,
	)
}
