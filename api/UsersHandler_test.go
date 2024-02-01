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

func TestUsersHandler_Post(t *testing.T) {
	// Create a new instance of UsersHandler with the mock repository
	r := repo.NewDummyRepository()
	handler := NewUsersHandler(r)

	// Create a sample eventUser to send in the request body

	user1 := models.GenerateValidUser()

	eventUserJSON, err := json.Marshal(user1)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, "/api/Users", bytes.NewBuffer(eventUserJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the PostsHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}
