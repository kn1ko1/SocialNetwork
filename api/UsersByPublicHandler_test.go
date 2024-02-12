package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"testing"
)

func TestUsersByPublicHandlerValidUserExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
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
}

func TestUsersByPublicHandlerInvalidUserExpectFail_Get(t *testing.T) {
	// Create a new instance of UsersHandler with the mock repository
	for i := 0; i < 10; i++ {
		handler := NewUsersHandler(R)

		user1 := models.GenerateInvalidUser()

		// Validate the user - it should fail due to the invalid email
		validateErr := user1.Validate()
		if validateErr == nil {
			t.Error("Expected invalid user to fail validation, if you see this the test has succeeded")
			return
		}

		eventUserJSON, err := json.Marshal(user1)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(eventUserJSON))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
			return
		}

		// Parse the response body to check if the error message is returned
		var errorMessage map[string]string
		err = json.Unmarshal(recorder.Body.Bytes(), &errorMessage)
		if err == nil {
			t.Fatal("Error decoding response body:", err)
		}

		// Check if the expected error message is returned
		expectedErrorMessage := "Invalid user data"
		if errorMessage["error"] == expectedErrorMessage {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, errorMessage["error"])
		}
	}
}
