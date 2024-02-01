package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
)

// Create a new instance of commentByIdHandler with the mock repository

// MockRepository is a mock implementation of IRepository for testing purposes
type MockCommentByIdHandlerRepository struct {
	Comment models.Comment
}

func TestCommentByIdHandler_Get(t *testing.T) {
	handler := NewCommentByIdHandler(R)
	comment, _ := handler.Repo.GetCommentById(1)

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

func TestCommentByIdHandler_Put(t *testing.T) {
	handler := NewCommentByIdHandler(R)
	comment, _ := handler.Repo.UpdateComment(*CommentExample)

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

func TestCommentByIdHandler_Delete(t *testing.T) {

	handler := NewCommentByIdHandler(R)

	commentJSON, err := json.Marshal(CommentExample)
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
