package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

/*
 author: wubo
*/

func FavoriteAction(c *gin.Context) {
	params := request.FavoriteRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	//获取当前请求的用户id
	svc := service.New(c)
	userId := c.GetUint("UserID")
	data, err := svc.FavoriteAction(params, userId)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}

func FavoriteListAction(c *gin.Context) {
	params := request.FavoriteListRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	svc := service.New(c)
	data, err := svc.FavoriteListAction(params)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}
