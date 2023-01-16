package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//Identity      string     `json:"identity" gorm:"index:,unique;type:varchar(36);not null;comment:用户唯一标识"`
	Username      string     `json:"username" gorm:"index:,unique;type:varchar(16);not null;comment:用户名"`
	Password      string     `json:"password" gorm:"type:varchar(64);not null;comment:密码"`
	Nickname      string     `json:"nickname" gorm:"type:varchar(20);comment:用户昵称"`
	HeadImg       string     `json:"head_img" gorm:"type:varchar(128);comment:用户头像"`
	FollowCount   int64      `json:"follow_count" gorm:"comment:关注总数"`
	FollowerCount int64      `json:"follower_count" gorm:"comment:粉丝总数"`
	Status        int8       `json:"status" gorm:"type:tinyint(8);default:1;comment:用户状态，1正常 2封号 3注销 4违规"`
	Videos        []Video    `json:"video" gorm:"foreignKey:UserID;references:ID;comment:视频列表"`
	Comments      []Comment  `json:"comments" gorm:"foreignKey:UserID;references:ID;comment:评论列表"`
	Favorites     []Favorite `json:"favorites" gorm:"foreignKey:UserID;references:ID;comment:喜欢列表"`
}

func (User) Table() string {
	return "user"
}
