package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	VideoID uint   `json:"video_id"`
	Content string `json:"content" gorm:"type:varchar(512);not null;comment:评论内容"`
}

func (Comment) TableName() string {
	return "comment"
}
