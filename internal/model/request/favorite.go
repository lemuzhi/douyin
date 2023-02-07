package request

/*
   author: wubo
   互动接口 请求参数结构体
*/

// FavoriteRequest  点赞操作
type FavoriteRequest struct {
	Token      string `json:"token" binding:"required"`
	VideoId    int64  `json:"video_id " binding:"required"`
	ActionType int32  `json:"action_type " binding:"required"`
}

// FavoriteListRequest 获取所有点赞视频
type FavoriteListRequest struct {
	Token  string `json:"token" binding:"required"`
	UserId int64  `json:"user_id " binding:"required"`
}

// CommentRequest  视频评论操作
type CommentRequest struct {
	Token       string `json:"token" binding:"required"`
	VideoId     int64  `json:"video_id " binding:"required"`
	ActionType  int32  `json:"action_type " binding:"required"`
	CommentText string `json:"comment_text" binding:"omitempty"`
	CommentId   int64  `json:"comment_id " binding:"omitempty"`
}

// CommentListRequest  获取视频的所有评论
type CommentListRequest struct {
	Token   string `json:"token" binding:"required"`
	VideoId int64  `json:"video_id " binding:"required"`
}
