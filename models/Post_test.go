package models

import (
	"fmt"
	"math/rand"
	"testing"
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
				t.Error("expect error for invalid post field")
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
				t.Error("expect error for missing post field")
			}
		})
	}
}

func TestPostValidateValidExpectNil(t *testing.T) {
	var posts []*Post
	for i := 0; i < tableRunCount; i++ {
		posts = append(posts, GenerateValidPost())
	}
	for _, p := range posts {
		name := fmt.Sprintf("%+v", *p)
		t.Run(name, func(t *testing.T) {
			err := p.Validate()
			if err != nil {
				t.Error("expect nil for valid post")
			}
		})
	}
}

func generateMissingFieldPost() *Post {
	p := GenerateValidPost()
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
	p := GenerateValidPost()
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
