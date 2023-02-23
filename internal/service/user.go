package service

import (
	"douyin/internal/model"
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"douyin/pkg/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (svc *Service) Register(params request.RegisterReq) (*response.LoginResponse, error) {
	pwd, err := utils.EncipherPassword(params.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username: params.Username,
		Password: pwd,
		//客户端用户注册暂无设置头像、背景图、简介功能，也无修改功能，所以暂时写死
		Avatar:          "https://c-ssl.duitang.com/uploads/blog/202102/08/20210208200511_45cb8.jpg",
		BackgroundImage: "https://article.autotimes.com.cn/wp-content/uploads/2022/04/95f35f8c40454bf1b4f18d7c79b5b948.jpg",
		Signature:       fmt.Sprintf("我于%s注册了抖声，欢迎关注！", time.Now().Format("2006-01-02 15:04:05")),
	}
	id, err := svc.dao.Register(&user)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(id)
	if err != nil {
		return nil, err
	}

	data := &response.LoginResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserID:   id,
		Token:    token,
	}
	return data, nil
}

func (svc *Service) Login(params request.LoginReq) (*response.LoginResponse, error) {

	user, err := svc.dao.Login(params.Username)
	if err != nil {
		return nil, err
	}

	//验证密码
	if !utils.VerifyPassword(user.Password, params.Password) {
		return nil, fmt.Errorf("%v", errcode.ErrUserOrPwd)
	}

	//生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	data := &response.LoginResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserID:   user.ID,
		Token:    token,
	}
	return data, nil
}

func (svc *Service) GetUserInfo(params request.UserReq) *response.UserInfoResponse {
	user, err := svc.dao.GetUserInfo(params.UserID)

	workCount, videoIdList := svc.dao.WorkCount(params.UserID) //作品数量
	favoriteCount := svc.dao.UserFavoriteCount(params.UserID)  //点赞数量
	totalFavorited := svc.dao.TotalFavorited(videoIdList)

	data := &response.UserInfoResponse{
		User: response.User{
			ID:              user.ID,
			Name:            user.Username,
			FollowCount:     svc.dao.FollowCount(user.ID),
			FollowerCount:   svc.dao.FollowerCount(user.ID),
			IsFollow:        false,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  totalFavorited,
			WorkCount:       workCount,
			FavoriteCount:   favoriteCount,
		},
	}

	if err != nil {
		//没有找到用户，用户不存在
		if err == gorm.ErrRecordNotFound {
			data.Response = errcode.NewResponse(errcode.ErrUserNotExist)
			return data
		}
		//其他错误
		data.Response = errcode.NewResponse(errcode.ErrService)
		return data
	}

	data.Response = errcode.NewResponse(errcode.OK)
	return data
}
