package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (svc *Service) RelationAction(c *gin.Context, params request.RelationActionReq) error {
	return svc.dao.RelationAction(c.GetUint("UserID"), params.ToUserID, uint8(params.ActionType))
}

func (svc *Service) FollowList(params request.FollowListReq) (*response.FollowListResponse, error) {
	//获取传入用户id关注的用户列表
	userIdList, err := svc.dao.GetFollowList(params.UserID)
	if err != nil {
		return nil, err
	}
	//存储所有我关注的用户id
	idList := make([]uint, len(userIdList))
	for i := 0; i < len(userIdList); i++ {
		idList[i] = userIdList[i].BeUserID
	}
	userMap, err := svc.dao.GetUserInfoList(idList)
	if err != nil {
		fmt.Println("FollowerList GetUserInfoList error: ", err)
		return nil, err
	}
	//获取当前用户的id
	myUid := svc.ctx.GetUint("UserID")

	var userList []*response.User
	for i := 0; i < len(idList); i++ {
		userMap[idList[i]].IsFollow, _ = svc.dao.IsFollow(myUid, userMap[idList[i]].ID)
		userList = append(userList, userMap[idList[i]])
	}

	data := response.FollowListResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserList: userList,
	}

	return &data, nil
}

func (svc *Service) FollowerList(params request.FollowListReq) (*response.FollowListResponse, error) {
	//获取关注传入用户id所属的粉丝列表
	userIdList, err := svc.dao.GetFollowerList(params.UserID)
	if err != nil {
		return nil, err
	}
	//存储传入用户关注的用户的id
	idList := make([]uint, len(userIdList))
	for i := 0; i < len(userIdList); i++ {
		idList[i] = userIdList[i].UserID
	}
	userMap, err := svc.dao.GetUserInfoList(idList)
	if err != nil {
		fmt.Println("FollowerList GetUserInfoList error: ", err)
		return nil, err
	}
	//获取当前用户的id
	myUid := svc.ctx.GetUint("UserID")

	var userList []*response.User
	for i := 0; i < len(idList); i++ {
		userMap[idList[i]].IsFollow, _ = svc.dao.IsFollow(myUid, userMap[idList[i]].ID)
		userList = append(userList, userMap[idList[i]])
	}

	data := response.FollowListResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserList: userList,
	}
	return &data, nil
}

func (svc *Service) FriendList(params request.FriendListReq) (*response.FriendListResponse, error) {
	userList, err := svc.dao.GetFriendList(params.UserID)
	if err != nil {
		return nil, err
	}

	data := response.FriendListResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserList: userList,
	}
	return &data, nil
}
