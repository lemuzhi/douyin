package request

type MessageRequest struct {
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserID   uint   `form:"to_user_id" json:"to_user_id " binding:"required"`
	ActionType int32  `form:"action_type" json:"action_type " binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
}

type MessageListRequest struct {
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserID   uint   `form:"to_user_id" json:"to_user_id " binding:"required"`
	PreMsgTime int64  `form:"pre_msg_time" json:"pre_msg_time " binding:"required"`
}
