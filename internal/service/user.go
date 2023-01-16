package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	errcode2 "douyin/pkg/errcode"
	"douyin/pkg/utils"
	"fmt"
	"gorm.io/gorm"
)

func (svc *Service) Login(params request.LoginReq) (*response.LoginResponse, error) {

	user, err := svc.dao.Login(params.Username)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(user.Password, params.Password) {
		return nil, fmt.Errorf("%v", errcode2.ErrUserOrPwd)
	}

	token, err := utils.GenerateToken(params.Username)
	if err != nil {
		return nil, err
	}
	//pwd, err := utils.EncipherPassword(params.Password)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("密码：", pwd)

	data := &response.LoginResponse{
		Response: errcode2.NewResponse(errcode2.OK),
		UserID:   1,
		Token:    token,
	}
	return data, nil
}

func (svc *Service) GetUserInfo(params request.UserReq) *response.UserInfoResponse {
	user, err := svc.dao.GetUserInfo(params.UserID)

	data := &response.UserInfoResponse{
		User: response.User{
			ID:            user.ID,
			Name:          user.Nickname,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		},
	}

	if err != nil {
		//没有找到用户，用户不存在
		if err == gorm.ErrRecordNotFound {
			data.Response = errcode2.NewResponse(errcode2.ErrUserNotExist)
			return data
		}
		//其他错误
		data.Response = errcode2.NewResponse(errcode2.ErrService)
		return data
	}

	data.Response = errcode2.NewResponse(errcode2.OK)
	return data
}
