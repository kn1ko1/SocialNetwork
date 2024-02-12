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

func TestUserPostsHandlerExpectPass_Get(t *testing.T) {
	handler := NewUserPostsHandler(R)
	post, _ := handler.Repo.GetPostsByUserId(1)

	userJSON, err := json.Marshal(post)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/users/1"

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

func TestUserPostsHandler_Get_Error(t *testing.T) {
	handler := NewUserPostsHandler(R)
	// Simulate an error from the repository
	_, err := handler.Repo.GetPostsByUserId(999) // Invalid user ID to trigger an error
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestUserPostsHandler_Put(t *testing.T) {
	handler := NewUserPostsHandler(R)
	user, _ := handler.Repo.UpdateUser(*models.GenerateValidUser())

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/users/" + fmt.Sprint(user.UserId)

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

func TestUserPostsHandler_Put_InvalidUserID(t *testing.T) {
	handler := NewUserPostsHandler(R)
	// Try to update a user with an invalid user ID
	_, err := handler.Repo.UpdateUser(*models.GenerateValidUser())
	if err == nil {
		t.Error("Expected error for invalid user ID, but got nil")
	}
}
