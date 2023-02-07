package dao

import (
	"douyin/internal/model"
	"gorm.io/gorm"
	"time"
)

const limit = 5 //每次获取的视频数量

func (dao *Dao) Feed() (video model.Video, user model.User, err error) {
	err = dao.db.Preload("User").Where("id=?", 1).First(&video).Error
	if err != nil {
		return
	}
	err = dao.db.Where("id=?", video.UserID).First(&user).Error
	if err != nil {
		return
	}
	return
}

func (dao *Dao) GetFeedList(lastTime time.Time) (videoList *[]*model.Video, err error) {
	err = dao.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "follow_count", "follower_count")
	}).Where("created_at < ?", lastTime).Order("created_at DESC").Limit(limit).Find(&videoList).Error

	return videoList, err
}
