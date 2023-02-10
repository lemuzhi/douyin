package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
	"log"
)

// PublishAction 投稿接口
func PublishAction(c *gin.Context) {
	send := errcode.New(c)
	// 单文件
	file, err := c.FormFile("data")
	if err != nil {
		log.Printf("get form err: %s", err.Error())
		send.RespFail(errcode.ErrPublishVideo)
		return
	}

	svc := service.New(c)
	err = svc.PublishAction(c, file)
	if err != nil {
		log.Printf("upload file err: %s", err.Error())
		send.RespFailDetail(errcode.ErrPublishVideo, err.Error())

		return
	}

	send.RespOK()
}

// PublishList 发布列表
func PublishList(c *gin.Context) {
	params := request.PublishListRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	data, err := svc.PublishList(params.UserID)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}
