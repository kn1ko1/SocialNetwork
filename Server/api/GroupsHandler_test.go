package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/Server/models"
	"testing"
)

func TestGroupsHandlerValidGroupExpectedPass_Post(t *testing.T) {
	handler := NewGroupsHandler(R)
	group := models.GenerateValidGroup()

	groupJSON, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groups"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(groupJSON))
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
}

func TestGroupsHandlerInValidGroupExpectedPass_Post(t *testing.T) {
	handler := NewGroupsHandler(R)
	group := models.GenerateInvalidGroup()

	groupJSON, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groups"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(groupJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestGroupsHandlerValidExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewGroupsHandler(R)

		URL := "/api/groups"

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodGet, URL, nil)
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
	}
}

func TestGroupsHandlerInvalidMethodExpectPass_Put1(t *testing.T) {
	handler := NewGroupsHandler(R)

	URL := "/api/groups"

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPut, URL, nil)
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
}

func TestGroupsHandlerInValidGroupExpectedPass_Put2(t *testing.T) {
	handler := NewGroupsHandler(R)
	group := models.GenerateInvalidGroup()

	groupJSON, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groups"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(groupJSON))
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
}
