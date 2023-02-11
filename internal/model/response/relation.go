package response

type RelationActionResponse struct {
	Response
}

type FollowListResponse struct {
	Response
	UserList []*User `json:"user_list"`
}
