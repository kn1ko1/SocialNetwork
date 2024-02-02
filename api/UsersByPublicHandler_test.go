package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersByPublicHandler_Get(t *testing.T) {
	handler := NewUsersByPublicHandler(R)
	users, _ := handler.Repo.GetUsersByPublic()

	userJSON, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/users/privacy/public"

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
