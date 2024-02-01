package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	apiTestVars "socialnetwork/api/apiTestVars"
	"socialnetwork/models"
	"socialnetwork/repo"
	"testing"
)

// MockEventUserRepository is a mock implementation of IRepository for testing purposes
// type MockEventUserRepository struct {
// 	EventUsers []models.EventUser
// }

func TestEventUsersHandler_Post(t *testing.T) {
	// Create a new instance of EventUsersHandler with the mock repository
	r := repo.NewDummyRepository()
	handler := NewEventUsersHandler(r)

	// Create a sample eventUser to send in the request body

	eventUser1 := models.EventUser{
		EventUserId: 1,
		CreatedAt:   apiTestVars.Timestamp,
		EventId:     1,
		UpdatedAt:   apiTestVars.Timestamp,
		UserId:      1,
	}

	eventUserJSON, err := json.Marshal(eventUser1)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, "/api/eventUsers", bytes.NewBuffer(eventUserJSON))
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

func TestEventUserHandler_Get(t *testing.T) {
	// Create a new instance of PostsHandler with the mock repository
	r := repo.NewDummyRepository()
	handler := NewEventUsersHandler(r)

	// Create a new HTTP request for a GET to "/api/eventUsers"
	req, err := http.NewRequest(http.MethodGet, "/api/eventUsers", nil)
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
