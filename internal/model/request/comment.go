package request

// CommentRequest  视频评论操作
type CommentRequest struct {
	Token       string `form:"token" json:"token" binding:"required"`
	VideoId     int64  `form:"video_id" json:"video_id " binding:"required"`
	ActionType  int32  `form:"action_type" json:"action_type " binding:"required"`
	CommentText string `form:"comment_text" json:"comment_text" binding:"omitempty"`
	CommentId   int64  `form:"comment_id" json:"comment_id " binding:"omitempty"`
}

// CommentListRequest  获取视频的所有评论
type CommentListRequest struct {
	Token   string `form:"token" json:"token" binding:"required"`
	VideoId int64  `form:"video_id" json:"video_id " binding:"required"`
}
