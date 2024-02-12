package models

import (
	"fmt"
	"testing"
)

func TestPostValidateInvalidFieldExpectError(t *testing.T) {
	var posts []*Post
	for i := 0; i < tableRunCount; i++ {
		posts = append(posts, GenerateInvalidPost())
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
		posts = append(posts, GenerateMissingFieldPost())
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
