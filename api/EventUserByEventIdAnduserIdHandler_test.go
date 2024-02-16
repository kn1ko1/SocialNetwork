package api

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestEventUserByEventIdAnduserIdHandlerValidMethodExpectedPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUserByEventIdAndUserIdHandler(R)
		eventId := rand.Intn(101)
		eventIdStr := strconv.Itoa(eventId)
		userId := rand.Intn(101)
		userIdStr := strconv.Itoa(userId)

		URL := "/api/events/" + eventIdStr + "/eventUsers/users/" + userIdStr

		// Create a new HTTP request with the encoded JSON as the request body
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
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventUserByEventIdAnduserIdHandlerInvalidMethodExpectedPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUserByEventIdAndUserIdHandler(R)
		eventId := rand.Intn(101)
		eventIdStr := strconv.Itoa(eventId)
		userId := rand.Intn(101)
		userIdStr := strconv.Itoa(userId)

		URL := "/api/events/" + eventIdStr + "/eventUsers/users/" + userIdStr

		// Create a new HTTP request with the encoded JSON as the request body
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
		// Add additional assertions as needed for your specific use case
	}
}
