package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroupUserByGroupIdAndUserIdHandler_Delete(t *testing.T) {

	handler := NewGroupUserByGroupIdAndUserIdHandler(R)
	err := handler.Repo.DeleteGroupUserByGroupIdAndUserId(GroupExample.GroupId, GroupExample.CreatorId)

	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groups"

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
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}
