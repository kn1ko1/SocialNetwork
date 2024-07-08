package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/Server/models"
)

func TestPostsHandlerValidPostExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewPostsHandler(R)
		post := models.GenerateValidPost()

		postJSON, err := json.Marshal(post)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(postJSON))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "multipart/form-data")

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the PostsHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	}
}

func TestPostsHandlerInValidPostExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewPostsHandler(R)
		post := models.GenerateInvalidPost()

		postJSON, err := json.Marshal(post)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(postJSON))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "multipart/form-data")

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the PostsHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	}
}

/**
func TestPostsHandlerExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewPostsHandler(R)

		// Create a new HTTP request for a GET to "/api/posts"
		req, err := http.NewRequest(http.MethodGet, "/api/posts", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the PostsHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}

		// Add additional assertions as needed for your specific use case
	}
}

**/
