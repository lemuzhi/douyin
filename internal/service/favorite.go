package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
)

func (svc *Service) FavoriteAction(params request.FavoriteRequest, userId uint) (response.FavoriteResponse, error) {

	response := response.FavoriteResponse{
		StatusCode: errcode.OK.Code,
		StatusMsg:  errcode.OK.Msg,
	}

	err := svc.dao.FavoriteAction(userId, params.VideoId, uint8(params.ActionType))
	return response, err
}

func (svc *Service) FavoriteListAction(params request.FavoriteListRequest) (response.FavoriteListResponse, error) {

	videos, err := svc.dao.FavoriteListAction(params.UserId)

	var videosRsp []response.VideoResponse
	//初始化返回参数
	for i := 0; i < len(videos); i++ {
		// 获取视频的作者信息
		user, err1 := svc.dao.FindUserByID(uint(videos[i].UserID))

		if err1 != nil {
			return response.FavoriteListResponse{}, err1
		}

		//查看当前用户是否关注了此用户isFollow
		followFlag, err2 := svc.dao.IsFollow(params.UserId, videos[i].UserID)
		if err2 != nil {
			return response.FavoriteListResponse{}, err2
		}

		userRsp := response.User{
			ID:            user.ID,
			Name:          user.Username,
			FollowCount:   svc.dao.FollowCount(user.ID),
			FollowerCount: svc.dao.FollowerCount(user.ID),
			IsFollow:      followFlag,
		}
		vRep := response.VideoResponse{
			ID:            videos[i].ID,
			Author:        userRsp,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: svc.dao.FavoriteCount(videos[i].ID),
			CommentCount:  svc.dao.CommentCount(videos[i].ID),
			IsFavorite:    true,
			Title:         videos[i].Title,
		}
		videosRsp = append(videosRsp, vRep)
	}

	return response.FavoriteListResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videosRsp,
	}, err
}
