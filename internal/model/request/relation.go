package request

type RelationActionReq struct {
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserID   uint   `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int32  `form:"action_type" json:"action_type" binding:"required"`
}

// TODO: 这种 UserID + Token的request 封装一下？

type CommonReq struct {
	UserID uint   `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

type FollowListReq struct {
	UserID uint   `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

type FriendListReq struct {
	UserID uint   `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}
