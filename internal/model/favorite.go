package model

import "time"

type Favorite struct {
	ID        uint  `gorm:"primarykey"`
	UserID    uint  `json:"user_id"`
	VideoID   uint  `json:"video_id"`
	Status    uint8 `json:"status" gorm:"index;type:tinyint(3);not null;comment:1已点赞 2已取消"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Favorite) TableName() string {
	return "favorite"
}
