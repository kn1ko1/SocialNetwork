package models

type Comment struct {
	CommentId int
	Body      string
	CreatedAt int64
	ImageURL  string
	PostId    int
	UpdatedAt int64
	UserId    int
}
