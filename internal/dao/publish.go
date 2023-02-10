package dao

import (
	"douyin/internal/model"
	"douyin/internal/model/response"
	"fmt"
)

func (dao *Dao) PublishAction(userID int64, title, playUrl, coverUrl string) error {
	video := model.Video{
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		UserID:   userID,
	}
	return dao.db.Create(&video).Error
}

func (dao *Dao) GetPublishList(userID string) (videoList []*response.VideoResponse, err error) {
	rows, err := dao.db.Model(&model.Video{}).Preload("User").Where("user_id = ?", userID).Rows()
	if err != nil {
		fmt.Println("GetPublishList Rows() error: ", err)
		return nil, err
	}

	defer rows.Close()

	var user model.User
	err = dao.db.Select("id", "username", "follow_count", "follower_count").Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var video model.Video
		// ScanRows 方法用于将一行记录扫描至结构体
		err = dao.db.ScanRows(rows, &video)
		if err != nil {
			fmt.Println("dao.db.ScanRows error: ", err)
		}

		// 业务逻辑...
		videoList = append(videoList, &response.VideoResponse{
			ID:    video.ID,
			Title: video.Title,
			Author: response.User{
				ID:            user.ID,
				Name:          user.Username,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		})
	}

	return videoList, nil
}
