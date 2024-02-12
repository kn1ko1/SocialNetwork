package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNewGroupUserByIdHandler_Get(t *testing.T) {
	groupUserId := rand.Intn(101)
	groupUserIdStr := strconv.Itoa(groupUserId)

	handler := NewGroupUserByIdHandler(R)

	URL := "/api/groupUsers/" + groupUserIdStr

	// Create a new HTTP request without a request body
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// You may also want to check the response body here if necessary
}

func TestNewGroupUserByIdHandler_Delete(t *testing.T) {
	handler := NewGroupUserByIdHandler(R)
	group, _ := handler.Repo.GetGroupUser(1)
	err := handler.Repo.DeleteGroupUser(group.GroupId)

	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groupUsers/" + fmt.Sprint(group.UserId)

	// Create a new HTTP request without a request body
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
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// You may also want to check the response body here if necessary
}
