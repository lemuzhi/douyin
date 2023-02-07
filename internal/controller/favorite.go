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

// FavoriteAction 点赞操作接口
func FavoriteAction(c *gin.Context) {
	params := request.FavoriteRequest{}
	send := errcode.New(c)

	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	data, err := svc.FavoriteAction(params)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}
