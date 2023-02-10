package request

type PublishActionRequest struct {
	Data  string `form:"data" binding:"required"`
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
}

type PublishListRequest struct {
	Token  string `form:"token" binding:"required"`
	UserID string `form:"user_id" binding:"required"`
}
