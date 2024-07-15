package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/Server/models"
	"testing"
)

func TestNotificationsHandlerValidNotificationExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewNotificationsHandler(R)
		notification := models.GenerateValidNotification()

		notifcationJSON, err := json.Marshal(notification)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(notifcationJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the NotificationsHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	}
}

func TestNotificationsHandlerInValidNotificationExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewNotificationsHandler(R)
		notification := models.GenerateInvalidNotification()

		notifcationJSON, err := json.Marshal(notification)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(notifcationJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the NotificationsHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	}
}
