package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
	"socialnetwork/repo"
)

func TestUsersHandlerValidUserExpectPass_Post(t *testing.T) {
	// Create a new instance of UsersHandler with the mock repository
	r := repo.NewDummyRepository()
	handler := NewUsersHandler(r)

	// Create a sample eventUser to send in the request body
	user1 := models.GenerateValidUser()

	eventUserJSON, err := json.Marshal(user1)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, "/api/Users", bytes.NewBuffer(eventUserJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the PostsHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		return
	}

	// Parse the response body to check if the user object is returned without an image URL
	var createdUser models.User
	err = json.Unmarshal(recorder.Body.Bytes(), &createdUser)
	if err != nil {
		t.Fatal("Error decoding response body:", err)
	}

}

func TestUsersHandlerInvalidUserExpectFail_Post(t *testing.T) {
	// Create a new instance of UsersHandler with the mock repository
	r := repo.NewDummyRepository()
	handler := NewUsersHandler(r)

	// Create a sample user to send in the request body
	user1 := models.User{
		UserId:    0,
		Bio:       "",
		CreatedAt: 6,
		DOB:       1,
		Email:     "hello", // This email is invalid
		FirstName: "",
		LastName:  "",
		UpdatedAt: 1,
		Username:  "",
	}

	// Validate the user - it should fail due to the invalid email
	validateErr := user1.Validate()
	if validateErr != nil {
		t.Error("Expected invalid user to fail validation, if you see this the test has succeeded")
		return
	} else {
		t.Error("Expected user to fail validation, but it passed")
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
	if err != nil {
		t.Fatal("Error decoding response body:", err)
	}

	// Check if the expected error message is returned
	expectedErrorMessage := "Invalid user data"
	if errorMessage["error"] != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, errorMessage["error"])
	}
}
