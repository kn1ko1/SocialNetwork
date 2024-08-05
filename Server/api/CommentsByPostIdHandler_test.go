package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommentsByPostIdHandlerValidPostIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		// Create a new instance of commentByIdHandler with the mock repository
		handler := NewCommentsByPostIdHandler(R)
		postId := RandomNumberStr

		URL := "/api/posts/" + postId + "/comments"

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

func TestCommentsByPostIdHandlerValidPostIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewCommentsByPostIdHandler(R)
		postId := RandomNumberStr

		URL := "/api/posts/" + postId + "/comments"

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

func TestCommentsByPostIdHandlerInValidMethodExpectPass_Post(t *testing.T) {
	handler := NewCommentsByPostIdHandler(R)
	postId := RandomNumberStr

	URL := "/api/posts/" + postId + "/comments"

	// Create a new HTTP request
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

func TestCommentsByPostIdHandlerInValidMethodExpectPass_Put(t *testing.T) {
	handler := NewCommentsByPostIdHandler(R)
	postId := RandomNumberStr

	URL := "/api/posts/" + postId + "/comments"

	// Create a new HTTP request
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
