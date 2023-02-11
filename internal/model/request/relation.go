package request

type RelationActionReq struct {
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserID   int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int32  `form:"action_type" json:"action_type" binding:"required"`
}

// type FollowListReq struct {
// 	UserID int64  `form:"user_id" json:"user_id" binding:"required"`
// 	Token  string `form:"token" json:"token" binding:"required"`
// }

type FollowListReq struct {
	Token  string `form:"token" json:"token" binding:"required"`
	UserID string `form:"user_id" json:"user_id" binding:"required"`
}
