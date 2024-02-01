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
type MockCommentByPostIdHandlerRepository struct {
	Comment models.Comment
}

func TestCommentByPostIdHandler_Get(t *testing.T) {

	// Create a new instance of commentByIdHandler with the mock repository
	handler := NewCommentsByPostIdHandler(R)
	comment, _ := handler.Repo.GetCommentById(1)

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, "/api/comments/1", bytes.NewBuffer(commentJSON))
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
func TestCommentByPostIdHandler_Put(t *testing.T) {
	handler := NewCommentsByPostIdHandler(R)
	comment, _ := handler.Repo.GetCommentById(1)

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPut, "/api/comments/1", bytes.NewBuffer(commentJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the commentByIdHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}

func TestCommentByPostIdHandler_Delete(t *testing.T) {
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
	req, err := http.NewRequest(http.MethodDelete, "/api/comments/1", bytes.NewBuffer(commentJSON))
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
