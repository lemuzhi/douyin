package request

type FeedRequest struct {
	LatestTime int64  `form:"latest_time" binding:"omitempty"`
	Token      string `form:"token" binding:"omitempty"`
}
