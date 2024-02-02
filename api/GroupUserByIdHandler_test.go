package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGroupUserByIdHandler_Get(t *testing.T) {

	handler := NewGroupUserByIdHandler(R)
	group, _ := handler.Repo.GetGroupUser(1)

	groupJSON, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groupUsers/" + fmt.Sprint(group.UserId)

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(groupJSON))
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
}

func TestNewGroupUserByIdHandler_Delete(t *testing.T) {

	handler := NewGroupUserByIdHandler(R)
	group, _ := handler.Repo.GetGroupUser(1)
	err := handler.Repo.DeleteGroupUser(group.GroupId)

	if err != nil {
		t.Fatal(err)
	}

	groupJSON, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groupUsers/" + fmt.Sprint(group.UserId)

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodDelete, URL, bytes.NewBuffer(groupJSON))
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
}
