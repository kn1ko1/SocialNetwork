package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventsByGroupIdHandlerValidGroupIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventsByGroupIdHandler(R)
		groupId := RandomNumberStr

		URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

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

func TestEventsByGroupIdHandlerValidGroupIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventsByGroupIdHandler(R)
		groupId := RandomNumberStr

		URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodDelete, URL, nil)
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

func TestEventsByGroupIdHandlerInValidMethodExpectPass_Post(t *testing.T) {
	handler := NewEventsByGroupIdHandler(R)
	groupId := RandomNumberStr

	URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, URL, nil)
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

func TestEventsByGroupIdHandlerInValidMethodExpectPass_Put(t *testing.T) {
	handler := NewEventsByGroupIdHandler(R)
	groupId := RandomNumberStr

	URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

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
