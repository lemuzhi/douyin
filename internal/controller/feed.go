package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"douyin/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFeedList(c *gin.Context) {
	params := request.FeedRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, errcode.NewResponse(errcode.ErrInvalidParams))
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	var uid uint
	// 判断用户是否登录，如登录，提取用户id
	if params.Token != "" {
		claims, _ := utils.ParseToken(params.Token)
		uid = claims.UserID
	}

	svc := service.New(c)
	data, err := svc.GetFeedList(uid, &params)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}
	fmt.Println("feed流", data)
	send.RespData(data)
}
