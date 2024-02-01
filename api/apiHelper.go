package api

import (
	"socialnetwork/models"
	"socialnetwork/repo"
	"time"
)

var CurrentTime = time.Now()

var Timestamp = CurrentTime.Unix()

var R = repo.NewDummyRepository()

var CommentExample = &models.Comment{
	CommentId: 1,
	Body:      "suicide squad",
	CreatedAt: Timestamp,
	ImageURL:  "imageurl",
	PostId:    1,
	UpdatedAt: Timestamp,
	UserId:    1,
}
