package service

import (
	"douyin/internal/dao"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New()
	return svc
}
