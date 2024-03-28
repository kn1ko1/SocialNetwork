package models

import (
	"errors"
	"math/rand"
)

type Post struct {
	PostId    int    `json:"postId"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"createdAt"`
	GroupId   int    `json:"groupId"`
	ImageURL  string `json:"imageURL"`
	Privacy   string `json:"privacy"`
	UpdatedAt int64  `json:"updatedAt"`
	UserId    int    `json:"userId"`
}

func (p *Post) Validate() error {
	// Validate logic here
	if p.Body == "" {
		return errors.New("invalid 'Body' field")
	}
	if p.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	// GroupID can be 0 - i.e. not posted to a Group - but cannot be negative
	if p.GroupId < -1 {
		return errors.New("invalid 'GroupId' field")
	}
	if p.Privacy != "public" && p.Privacy != "private" && p.Privacy != "almost private" {
		return errors.New("invalid 'Privacy' field")
	}
	if p.UpdatedAt < p.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field")
	}
	if p.UserId <= 0 {
		return errors.New("invalid 'UserId' field")
	}
	return nil
}

func GenerateValidPost() *Post {
	ctime := rand.Int63n(1000) + 1
	idxBody := rand.Intn(len(sutBody))
	idxImageURL := rand.Intn(len(sutImageURL))
	idxPrivacy := rand.Intn(len(sutPrivacy))

	p := &Post{
		Body:      sutBody[idxBody],
		CreatedAt: ctime,
		GroupId:   rand.Intn(1000) + 1,
		ImageURL:  sutImageURL[idxImageURL],
		Privacy:   sutPrivacy[idxPrivacy],
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return p
}

func GenerateMissingFieldPost() *Post {
	p := GenerateValidPost()
	missingField := rand.Intn(5)
	switch missingField {
	case 0:
		p.Body = ""
	case 1:
		p.CreatedAt = 0
	case 2:
		p.Privacy = ""
	case 3:
		p.UpdatedAt = 0
	case 4:
		p.UserId = 0
	}
	return p
}

func GenerateInvalidPost() *Post {
	p := GenerateValidPost()
	invalidField := rand.Intn(5)
	switch invalidField {
	case 0:
		p.Body = ""
	case 1:
		p.CreatedAt = -p.CreatedAt
	case 2:
		p.Privacy = ""
	case 3:
		p.UpdatedAt = -p.UpdatedAt
	case 4:
		p.UserId = -p.UserId
	}
	return p
}
