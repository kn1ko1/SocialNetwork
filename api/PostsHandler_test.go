package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
	"socialnetwork/repo"
)

func TestPostsHandler_Post(t *testing.T) {
	// Create a new instance of PostsHandler with the mock repository
	r := repo.NewDummyRepository()
	handler := NewPostsHandler(r)

	// Create a sample post to send in the request body
	post := models.Post{
		PostId:    1,
		Body:      "Test body",
		CreatedAt: Timestamp,
		GroupId:   1,
		ImageURL:  "poop",
		UpdatedAt: Timestamp,
		UserId:    1,
	}

	postJSON, err := json.Marshal(post)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, "/api/posts", bytes.NewBuffer(postJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the PostsHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}

func TestPostsHandler_Get(t *testing.T) {
	// Create a new instance of PostsHandler with the mock repository
	r := repo.NewDummyRepository()
	handler := NewPostsHandler(r)

	// Create a new HTTP request for a GET to "/api/posts"
	req, err := http.NewRequest(http.MethodGet, "/api/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the PostsHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// Add additional assertions as needed for your specific use case
}
