package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventUserByEventIdAnduserIdHandlerValidMethodExpectedPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUserByEventIdAndUserIdHandler(R)
		eventId := RandomNumberStr
		userId := RandomNumberStr

		URL := "/api/events/" + eventId + "/eventUsers/users/" + userId

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
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	}
}

func TestEventUserByEventIdAnduserIdHandlerInvalidMethodExpectedPass_Put(t *testing.T) {
	handler := NewEventUserByEventIdAndUserIdHandler(R)
	eventId := RandomNumberStr
	userId := RandomNumberStr

	URL := "/api/events/" + eventId + "/eventUsers/users/" + userId

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

func TestEventUserByEventIdAnduserIdHandlerInvalidMethodExpectedPass_Post(t *testing.T) {
	handler := NewEventUserByEventIdAndUserIdHandler(R)
	eventId := RandomNumberStr
	userId := RandomNumberStr

	URL := "/api/events/" + eventId + "/eventUsers/users/" + userId

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
