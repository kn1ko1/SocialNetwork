package transport

type UserTransport struct {
	UserId     int    `json:"userId"`
	IsFollowed bool   `json:"isFollowed"`
	IsMember   bool   `json:"isMember"`
	Username   string `json:"username"`
}
