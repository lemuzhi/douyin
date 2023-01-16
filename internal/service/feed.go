package service

import (
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"time"
)

func (svc *Service) Feed() (response.FeedResponse, error) {
	video, user, err := svc.dao.Feed()
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

	return response.FeedResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videoRsp,
		NextTime:  time.Now().Unix(),
	}, err
}
