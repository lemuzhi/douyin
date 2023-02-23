package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// MessageList 消息查询
func MessageList(c *gin.Context) {

	var params request.MessageListRequest
	send := errcode.New(c)
	err1 := c.ShouldBindQuery(&params)
	if err1 != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	UserId := c.GetUint("UserID")
	data, err2 := svc.MessageListAction(&params, UserId)
	if err2 != nil {
		send.RespFail(errcode.Fail)
		return
	}
	send.RespData(data)
}

// MessageAction 消息发送
func MessageAction(c *gin.Context) {

	var params request.MessageRequest
	send := errcode.New(c)
	if err := c.ShouldBindQuery(&params); err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	userId := c.GetUint("UserID")
	data, err := svc.MessageAction(&params, userId)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}
	send.RespData(data)
}
