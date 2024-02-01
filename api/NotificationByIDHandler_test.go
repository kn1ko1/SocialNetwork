package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"testing"
)

func TestNotificationByIdHandler_Get(t *testing.T) {
	handler := NewNotificationByIdHandler(R)
	posts, _ := handler.Repo.GetNotificationById(1)

	userJSON, err := json.Marshal(posts)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/posts/1"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(userJSON))
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

func TestNotificationByIdHandler_Put(t *testing.T) {
	handler := NewNotificationByIdHandler(R)
	notifcation, _ := handler.Repo.UpdateNotification(*models.GenerateValidNotification())
	notifcationJSON, err := json.Marshal(notifcation)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/comments/" + fmt.Sprint(notifcation.NotificationId)

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(notifcationJSON))
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
