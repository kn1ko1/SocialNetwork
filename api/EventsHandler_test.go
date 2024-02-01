package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventsHandler_Post(t *testing.T) {

	handler := NewEventsHandler(R)
	event, _ := handler.Repo.CreateEvent(*EventExample)

	eventJSON, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/events"

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

func TestEventsHandler_Get(t *testing.T) {
	handler := NewEventsHandler(R)
	events, _ := handler.Repo.GetAllEvents()

	eventJSON, err := json.Marshal(events)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/events"

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
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}
