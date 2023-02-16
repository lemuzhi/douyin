package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"log"

	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	params := request.RelationActionReq{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	// 不应该能关注自己
	UserID := c.GetUint("UserID")
	BeUserID := params.ToUserID
	if UserID == BeUserID {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	err = svc.RelationAction(c, params)

	if err != nil {
		log.Printf("relation action err: %s", err.Error())
		// TODO: errcode
		send.RespFailDetail(errcode.Fail, err.Error())

		return
	}

	send.RespOK()
}

func FollowList(c *gin.Context) {
	params := request.FollowListReq{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	data, err := svc.FollowList(params)
	// dousheng前端有bug，user_id始终为0，已有人反应

	if err != nil {
		log.Printf("get follow list err: %s", err.Error())
		send.RespFailDetail(errcode.Fail, err.Error())

		return
	}

	send.RespData(data)
}

func FollowerList(c *gin.Context) {
	params := request.FollowListReq{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	data, err := svc.FollowerList(params)

	if err != nil {
		log.Printf("get follow list err: %s", err.Error())
		send.RespFailDetail(errcode.Fail, err.Error())

		return
	}

	send.RespData(data)
}
