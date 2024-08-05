package models

import (
	"fmt"
	"testing"
)

const (
	tableRunCount = 10
)

func TestCommentValidateInvalidFieldExpectError(t *testing.T) {
	var comments []*Comment
	for i := 0; i < tableRunCount; i++ {
		comments = append(comments, GenerateInvalidComment())
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
		comments = append(comments, GenerateMissingFieldComment())
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
