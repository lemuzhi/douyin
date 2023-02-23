package response

type VideoResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        *User  `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}
