package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotificationByUserIdHandlerValidRequestExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewNotificationsByUserIdHandler(R)
		userId := RandomNumberStr

		URL := "/api/notifications/" + userId

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodGet, URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the Handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	}
}

func TestNotificationByUserIdHandlerValidRequestExpectPass_Post(t *testing.T) {
	handler := NewNotificationsByUserIdHandler(R)
	userId := RandomNumberStr

	URL := "/api/notifications/" + userId

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the Handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}

func TestNotificationByUserIdHandlerValidRequestExpectPass_Put(t *testing.T) {
	handler := NewNotificationsByUserIdHandler(R)
	userId := RandomNumberStr

	URL := "/api/notifications/" + userId

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPut, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the Handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}
