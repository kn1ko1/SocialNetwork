package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
)

// MockRepository is a mock implementation of IRepository for testing purposes
type MockCommentsHandlerRepository struct {
	Comment models.Comment
}

func TestCommentsHandler_Post(t *testing.T) {
	handler := NewCommentsHandler(R)
	comment, _ := handler.Repo.CreateComment(*CommentExample)

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/comments"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(commentJSON))
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

func TestCommentsHandler_Get(t *testing.T) {

	// Create a new instance of commentByIdHandler with the mock repository
	handler := NewCommentsHandler(R)
	comment, _ := handler.Repo.GetAllComments()

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/comments"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(commentJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the GetCommentsByPostId
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}
