package request

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required,min=6,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

type UserReq struct {
	UserID int64  `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}
