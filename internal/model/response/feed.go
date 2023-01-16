package response

type FeedResponse struct {
	Response
	NextTime  int64           `json:"next_time"`
	VideoList []VideoResponse `json:"video_list"`
}
