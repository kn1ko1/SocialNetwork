package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestValidateInvalidCommentExpectError(t *testing.T) {
	var vc []*Comment
	for i := 0; i < 10; i++ {
		vc = append(vc, generateInvalidComment())
	}
	for _, c := range vc {
		name := fmt.Sprintf("%+v", c)
		t.Run(name, func(t *testing.T) {
			err := c.Validate()
			if err == nil {
				t.Error("Expect error for invalid comment")
			}
		})
	}
}

func TestValidateValidCommentExpectNil(t *testing.T) {
	var vc []*Comment
	for i := 0; i < 10; i++ {
		vc = append(vc, generateValidComment())
	}
	for _, c := range vc {
		name := fmt.Sprintf("%+v", c)
		t.Run(name, func(t *testing.T) {
			err := c.Validate()
			if err != nil {
				t.Error("Expect nil for valid comment")
			}
		})
	}
}

func generateValidComment() *Comment {
	ctime := rand.Int63n(1000)
	c := &Comment{
		Body:      "hello",
		CreatedAt: ctime,
		PostId:    rand.Intn(1000),
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000),
	}
	return c
}

func generateInvalidComment() *Comment {
	c := generateValidComment()
	missingField := rand.Intn(5)
	switch missingField {
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
