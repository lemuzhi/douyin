package request

type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,min=6,max=32"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=32"`
}

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required,min=6,max=32"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=32"`
}

type UserReq struct {
	UserID uint   `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}
