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
type MockCommentsByPostIdHandlerRepository struct {
	Comment models.Comment
}

func TestCommentsByPostIdHandler_Get(t *testing.T) {

	// Create a new instance of commentByIdHandler with the mock repository
	handler := NewCommentsByPostIdHandler(R)
	comment, err := handler.Repo.GetCommentsByPostId(1)

	if err != nil {
		t.Fatal(err)
	}

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, "/api/posts/1/comment", bytes.NewBuffer(commentJSON))
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

func TestCommentsByPostIdHandler_Delete(t *testing.T) {
	handler := NewCommentsByPostIdHandler(R)
	// err := handler.Repo.DeleteCommentById(comment)
	dcomment := &models.Comment{
		CommentId: 1,
		Body:      "suicide squad",
		CreatedAt: Timestamp,
		ImageURL:  "imageurl",
		PostId:    1,
		UpdatedAt: Timestamp,
		UserId:    1,
	}

	commentJSON, err := json.Marshal(dcomment)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodDelete, "/api/posts/1/comments", bytes.NewBuffer(commentJSON))
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
