package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// CommentAction 评论操作
func CommentAction(c *gin.Context) {
	params := request.CommentRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	//获取当前请求的用户id
	svc := service.New(c)
	userId := c.GetInt64("UserID")
	data, err := svc.CommentAction(params, userId)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}

func CommentListAction(c *gin.Context) {
	params := request.CommentListRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	//获取当前请求的用户id
	svc := service.New(c)
	userId := c.GetInt64("UserID")
	data, err := svc.CommentListAction(params, userId)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}
