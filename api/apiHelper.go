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

var EventExample = &models.Event{
	EventId:     1,
	CreatedAt:   Timestamp,
	DateTime:    Timestamp,
	Description: "Event",
	GroupId:     1,
	Title:       "Event",
	UpdatedAt:   Timestamp,
	UserId:      1,
}

var EventUserExample = &models.EventUser{
	EventUserId: 1,
	CreatedAt:   Timestamp,
	EventId:     1,
	UpdatedAt:   Timestamp,
	UserId:      1,
}
