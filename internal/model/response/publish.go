package response

type PublishActionResponse struct {
	Data  string `form:"data" binding:"required"`
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
}

type PublishListResponse struct {
	Response
	VideoList []*VideoResponse `json:"video_list"`
}
