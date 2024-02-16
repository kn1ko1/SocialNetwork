package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGroupUserByIdHandlerValidIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewGroupUserByIdHandler(R)
		groupUserId := RandomNumberStr

		URL := "/api/groupUsers/" + groupUserId

		// Create a new HTTP request
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
	}
}

func TestNewGroupUserByIdHandlerValidIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewGroupUserByIdHandler(R)
		groupUserId := RandomNumberStr

		URL := "/api/groupUsers/" + groupUserId

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
}

func TestNewGroupUserByIdHandlerInValidMethodExpectPass_Post(t *testing.T) {
	handler := NewGroupUserByIdHandler(R)
	group, _ := handler.Repo.GetGroupUser(rand.Intn(101))
	err := handler.Repo.DeleteGroupUser(group.GroupId)

	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/groupUsers/" + fmt.Sprint(group.UserId)

	// Create a new HTTP request without a request body
	req, err := http.NewRequest(http.MethodPost, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}

func TestNewGroupUserByIdHandlerInValidMethodExpectPass_Put(t *testing.T) {
	handler := NewGroupUserByIdHandler(R)
	groupUserId := RandomNumberStr

	URL := "/api/groupUsers/" + groupUserId

	// Create a new HTTP request without a request body
	req, err := http.NewRequest(http.MethodPut, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}
