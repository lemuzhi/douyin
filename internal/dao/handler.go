package dao

import (
	"douyin/initialize"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao struct {
	db  *gorm.DB
	rdb *redis.Client
}

func New() *Dao {
	return &Dao{
		db:  initialize.DB,
		rdb: initialize.RDB,
	}
}
