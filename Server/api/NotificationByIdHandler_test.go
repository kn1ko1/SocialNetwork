package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/Server/models"
	"strconv"
	"testing"
)

func TestNotificationByIdHandlerValidRequestExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewNotificationByIdHandler(R)
		notificationId := RandomNumberStr

		URL := "/api/notifications/" + notificationId

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

func TestNotificationByIdHandlerValidRequestExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewNotificationByIdHandler(R)
		notification := models.GenerateValidNotification()
		notifcationJSON, err := json.Marshal(notification)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/notifications/" + strconv.Itoa(notification.NotificationId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(notifcationJSON))
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

func TestNotificationByIdHandlerValidRequestExpectPass_Post1(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewNotificationByIdHandler(R)
		notificationId := RandomNumberStr

		URL := "/api/notifications/" + notificationId

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
}

func TestNotificationByIdHandlerValidRequestExpectPass_Post2(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewNotificationByIdHandler(R)
		notification := models.GenerateValidNotification()
		notifcationJSON, err := json.Marshal(notification)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/notifications/" + strconv.Itoa(notification.NotificationId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(notifcationJSON))
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
}

func TestNotificationByIdHandlerInValidRequestExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewNotificationByIdHandler(R)
		notification := models.GenerateInvalidNotification()
		notifcationJSON, err := json.Marshal(notification)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/notifications/" + strconv.Itoa(notification.NotificationId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(notifcationJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the Handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	}
}

func TestNotificationByIdHandlerValidRequestExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewNotificationByIdHandler(R)
		notificationId := RandomNumberStr

		URL := "/api/notifications/" + notificationId

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodDelete, URL, nil)
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
