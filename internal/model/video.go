package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	//Identity      string     `json:"identity" gorm:"index:,unique;type:varchar(20);comment:视频唯一标识"`
	Title         string     `json:"title" gorm:"index;type:varchar(256);comment:视频标题"`
	PlayUrl       string     `json:"play_url" gorm:"type:varchar(128);comment:视频播放地址"`
	CoverUrl      string     `json:"cover_url" gorm:"type:varchar(128);comment:视频封面地址"`
	FavoriteCount int64      `json:"favorite_count" gorm:"comment:视频点赞总数"`
	CommentCount  int64      `json:"comment_count" gorm:"comment:视频评论总数"`
	Comments      []Comment  `gorm:"foreignKey:VideoID;references:ID;comment:评论列表"`
	Favorites     []Favorite `gorm:"foreignKey:VideoID;references:ID;comment:喜欢列表"`
	UserID        int64      `json:"user_id"`
	User          User
}

func (Video) TableName() string {
	return "video"
}
