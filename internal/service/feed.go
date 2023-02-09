package service

import (
	"douyin/internal/model"
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
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

	videoRsp := []*response.VideoResponse{
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

func (svc *Service) GetFeedList(c *gin.Context, params *request.FeedRequest) (resp response.FeedResponse, err error) {
	var lastTime time.Time
	if params.LatestTime == 0 {
		//如果未传入时间戳，使用当前时间
		lastTime = time.Now().Local()
	} else {
		//对传入的时间戳进行处理
		lastTime = time.UnixMilli(params.LatestTime)
	}
	userID := c.GetInt64("UserID")
	videoList, err := svc.dao.GetFeedList(lastTime)

	var isFollow, isFavorite bool
	var video *model.Video

	//以登录逻辑
	if userID > 0 {
		//TODO 判断用户用户是否点赞该视频，是否关注该视频的用户
	}

	var videos []*response.VideoResponse
	for i := 0; i < len(*videoList); i++ {
		video = (*videoList)[i]
		videos = append(videos, &response.VideoResponse{
			ID:    video.ID,
			Title: video.Title,
			Author: response.User{
				ID:            video.User.ID,
				Name:          video.User.Username,
				FollowCount:   video.User.FollowCount,
				FollowerCount: video.User.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
		})
	}

	return response.FeedResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videos,
		NextTime:  (*videoList)[len(*videoList)-1].CreatedAt.UnixMilli(),
	}, err
}
