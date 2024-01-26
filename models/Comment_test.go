package models

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	tableRunCount = 10
)

func TestCommentValidateInvalidFieldExpectError(t *testing.T) {
	var comments []*Comment
	for i := 0; i < tableRunCount; i++ {
		comments = append(comments, generateInvalidComment())
	}
	for _, c := range comments {
		name := fmt.Sprintf("%+v", *c)
		t.Run(name, func(t *testing.T) {
			err := c.Validate()
			if err == nil {
				t.Error("expect error for invalid comment field")
			}
		})
	}
}

func TestCommentValidateMissingFieldExpectError(t *testing.T) {
	var comments []*Comment
	for i := 0; i < tableRunCount; i++ {
		comments = append(comments, generateMissingFieldComment())
	}
	for _, c := range comments {
		name := fmt.Sprintf("%+v", *c)
		t.Run(name, func(t *testing.T) {
			err := c.Validate()
			if err == nil {
				t.Error("expect error for missing comment field")
			}
		})
	}
}

func TestCommentValidateValidExpectNil(t *testing.T) {
	var comments []*Comment
	for i := 0; i < tableRunCount; i++ {
		comments = append(comments, GenerateValidComment())
	}
	for _, c := range comments {
		name := fmt.Sprintf("%+v", *c)
		t.Run(name, func(t *testing.T) {
			err := c.Validate()
			if err != nil {
				t.Error("expect nil for valid comment")
			}
		})
	}
}

func GenerateValidComment() *Comment {
	ctime := rand.Int63n(1000) + 1
	idx := rand.Intn(len(sutBody))
	c := &Comment{
		Body:      sutBody[idx],
		CreatedAt: ctime,
		PostId:    rand.Intn(1000) + 1,
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return c
}

func generateMissingFieldComment() *Comment {
	c := GenerateValidComment()
	missingField := rand.Intn(5)
	switch missingField {
	case 0:
		c.Body = ""
	case 1:
		c.CreatedAt = 0
	case 2:
		c.PostId = 0
	case 3:
		c.UpdatedAt = 0
	case 4:
		c.UserId = 0
	}
	return c
}

func generateInvalidComment() *Comment {
	c := GenerateValidComment()
	invalidField := rand.Intn(5)
	switch invalidField {
	case 0:
		c.Body = ""
	case 1:
		c.CreatedAt = -c.CreatedAt
	case 2:
		c.PostId = -c.PostId
	case 3:
		c.UpdatedAt = -c.UpdatedAt
	case 4:
		c.UserId = -c.UserId
	}
	return c
}
