package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"testing"
)

func TestPostByIdHandler_Get(t *testing.T) {
	handler := NewPostByIdHandler(R)
	posts, _ := handler.Repo.GetPostById(1)

	userJSON, err := json.Marshal(posts)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/posts/1"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the Handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestPostByIdHandlerValidPostExpectPass_Put(t *testing.T) {
	handler := NewPostByIdHandler(R)
	comment, _ := handler.Repo.UpdateComment(*models.GenerateValidComment())

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/comments/" + fmt.Sprint(comment.CommentId)

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(commentJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the commentByIdHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}
