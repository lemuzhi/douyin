package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Feed(c *gin.Context) {
	params := request.FeedRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, errcode.NewResponse(errcode.ErrInvalidParams))
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	svc := service.New(c)
	data, err := svc.Feed()
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}
	send.RespData(data)
}
