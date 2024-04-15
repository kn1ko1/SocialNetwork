package transport

type UserTransport struct {
	UserId     int    `json:"userId"`
	IsFollowed bool   `json:"isFollowed"`
	Username   string `json:"username"`
}
