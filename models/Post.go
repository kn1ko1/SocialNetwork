package models

type Post struct {
	PostId    int
	Body      string
	CreatedAt int64
	GroupId   int
	ImageURL  string
	UpdatedAt int64
	UserId    int
}
