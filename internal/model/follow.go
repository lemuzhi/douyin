package model

import (
	"time"
)

type Follow struct {
	ID        uint  `gorm:"primarykey"`
	UserID    uint  `json:"user_id" gorm:"index;type:int;not null;comment:关注者用户ID"`     //关注者用户ID
	BeUserID  uint  `json:"be_user_id" gorm:"index;type:int;not null;comment:被关注者用户ID"` //被关在者用户ID
	Status    uint8 `json:"status" gorm:"index;type:tinyint(2);not null;comment:1已关注 2已取消"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Follow) TableName() string {
	return "follow"
}
