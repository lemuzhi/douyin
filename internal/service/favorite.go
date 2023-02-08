package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
)

func (svc *Service) FavoriteAction(params request.FavoriteRequest, userId int64) (response.FavoriteResponse, error) {

	response := response.FavoriteResponse{
		StatusCode: errcode.OK.Code,
		StatusMsg:  errcode.OK.Msg,
	}
	err := svc.dao.FavoriteAction(userId, params.VideoId, uint8(params.ActionType))
	return response, err
}

func (svc *Service) FavoriteListAction(params request.FavoriteListRequest) (response.FavoriteListResponse, error) {

	video, user, err := svc.dao.FavoriteListAction(params.UserId)
	userRsp := response.User{
		ID:            user.ID,
		Name:          user.Nickname,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	videoRsp := []response.VideoResponse{
		{
			ID:            video.ID,
			Title:         video.Title,
			Author:        userRsp,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		},
	}
	return response.FavoriteListResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videoRsp,
	}, err
}
