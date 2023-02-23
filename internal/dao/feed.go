package dao

import (
	"douyin/internal/model"
	"gorm.io/gorm"
	"time"
)

const limit = 5 //每次获取的视频数量

func (dao *Dao) GetFeedList(lastTime time.Time) (videoList *[]*model.Video, err error) {
	err = dao.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "avatar", "background_image", "signature")
	}).Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id", "video_id")
	}).Preload("Favorites", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id", "video_id").Where("status = ?", 1)
	}).Where("created_at < ?", lastTime).Order("created_at DESC").Limit(limit).Find(&videoList).Error

	return videoList, err
}
