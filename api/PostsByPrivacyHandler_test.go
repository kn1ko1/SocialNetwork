package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
)

// Create a new instance of userByIdHandler with the mock repository

// MockRepository is a mock implementation of IRepository for testing purposes
type MockPostsByPrivacyHandlerRepository struct {
	User models.User
}

func TestPostsByPrivacyHandlerExpectPass_Get(t *testing.T) {
	handler := NewPostsByPrivacyHandler(R)
	posts, _ := handler.Repo.GetPostsByPrivacy("public")

	userJSON, err := json.Marshal(posts)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/posts/privacy/public"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the userByIdHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}
