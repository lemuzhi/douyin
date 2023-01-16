package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  int64  `json:"user_id"`
	VideoID int64  `json:"video_id"`
	Content string `json:"content" gorm:"type:varchar(512);not null;comment:评论内容"`
}

func (Comment) TableName() string {
	return "comment"
}
