package request

// FavoriteRequest  点赞操作
type FavoriteRequest struct {
	Token      string `form:"token" json:"token" binding:"required"`
	VideoId    int64  `form:"video_id" json:"video_id " binding:"required"`
	ActionType int32  `form:"action_type" json:"action_type " binding:"required"`
}

// FavoriteListRequest 获取所有点赞视频
type FavoriteListRequest struct {
	Token  string `form:"token" json:"token" binding:"required"`
	UserId int64  `form:"user_id" json:"user_id " binding:"required"`
}
