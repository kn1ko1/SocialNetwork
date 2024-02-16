package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"testing"
)

func TestEventUsersHandlerValidEventUserExpectedPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUsersHandler(R)
		eventUser := models.GenerateValidEventUser()

		eventJSON, err := json.Marshal(eventUser)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/eventUsers"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventUsersHandlerInvalidMethodExpectedPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUsersHandler(R)
		eventUser := models.GenerateValidEventUser()

		eventJSON, err := json.Marshal(eventUser)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/eventUsers"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventUsersHandlerInvalidEventUserExpectedPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUsersHandler(R)
		eventUser := models.GenerateInvalidEventUser()

		eventJSON, err := json.Marshal(eventUser)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/eventUsers"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
		// Add additional assertions as needed for your specific use case
	}
}
