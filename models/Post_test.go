package models

import (
	"fmt"
	"math/rand"
	"testing"
)

var (
	sutImageURL = []string{"URL1", "URL2", "URL3"}
)

func TestPostValidateInvalidFieldExpectError(t *testing.T) {
	var posts []*Post
	for i := 0; i < tableRunCount; i++ {
		posts = append(posts, generateInvalidPost())
	}
	for _, p := range posts {
		name := fmt.Sprintf("%+v", *p)
		t.Run(name, func(t *testing.T) {
			err := p.Validate()
			if err == nil {
				t.Error("expect error for invalid event field")
			}
		})
	}
}

func TestPostValidateMissingFieldExpectError(t *testing.T) {
	var posts []*Post
	for i := 0; i < tableRunCount; i++ {
		posts = append(posts, generateMissingFieldPost())
	}
	for _, p := range posts {
		name := fmt.Sprintf("%+v", *p)
		t.Run(name, func(t *testing.T) {
			err := p.Validate()
			if err == nil {
				t.Error("expect error for missing event field")
			}
		})
	}
}

func TestPostValidateValidExpectNil(t *testing.T) {
	var posts []*Post
	for i := 0; i < tableRunCount; i++ {
		posts = append(posts, generateValidPost())
	}
	for _, p := range posts {
		name := fmt.Sprintf("%+v", *p)
		t.Run(name, func(t *testing.T) {
			err := p.Validate()
			if err != nil {
				t.Error("expect nil for valid event")
			}
		})
	}
}

func generateValidPost() *Post {
	ctime := rand.Int63n(1000) + 1
	idxBody := rand.Intn(len(sutBody))
	idxImageURL := rand.Intn(len(sutImageURL))

	p := &Post{
		Body:      sutBody[idxBody],
		CreatedAt: ctime,
		GroupId:   rand.Intn(1000) + 1,
		ImageURL:  sutImageURL[idxImageURL],
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return p
}

func generateMissingFieldPost() *Post {
	p := generateValidPost()
	missingField := rand.Intn(4)
	switch missingField {
	case 0:
		p.Body = ""
	case 1:
		p.CreatedAt = 0
	case 2:
		p.UpdatedAt = 0
	case 3:
		p.UserId = 0
	}
	return p
}

func generateInvalidPost() *Post {
	p := generateValidPost()
	invalidField := rand.Intn(4)
	switch invalidField {
	case 0:
		p.Body = ""
	case 1:
		p.CreatedAt = -p.CreatedAt
	case 2:
		p.UpdatedAt = -p.CreatedAt
	case 3:
		p.UserId = -p.UserId
	}
	return p
}
