package transport

type PostFromFrontend struct {
	PostId    int    `json:"postId"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"createdAt"`
	GroupId   int    `json:"groupId"`
	Privacy   string `json:"privacy"`
	UpdatedAt int64  `json:"updatedAt"`
	UserId    int    `json:"userId"`
}
