package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
)

func TestUserByIdHandler_Get(t *testing.T) {
	handler := NewUserByIdHandler(R)
	user, _ := handler.Repo.GetUserById(1)

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/users/" + fmt.Sprint(user.UserId)

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

func TestUserByIdHandler_Put(t *testing.T) {
	handler := NewUserByIdHandler(R)
	user, _ := handler.Repo.UpdateUser(*models.GenerateValidUser())

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/users/" + fmt.Sprint(user.UserId)

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
