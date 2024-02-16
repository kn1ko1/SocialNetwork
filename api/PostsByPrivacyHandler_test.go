package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostsByPrivacyHandlerValidPublicExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewPostsByPrivacyHandler(R)
		posts, _ := handler.Repo.GetPostsByPrivacy("public")

		userJSON, err := json.Marshal(posts)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts/privacy/public"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(userJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the userByIdHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	}
}

func TestPostsByPrivacyHandlerValidPrivateExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewPostsByPrivacyHandler(R)
		posts, _ := handler.Repo.GetPostsByPrivacy("private")

		userJSON, err := json.Marshal(posts)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts/privacy/private"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(userJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the userByIdHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	}
}

func TestPostsByPrivacyHandlerInValidPrivateExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewPostsByPrivacyHandler(R)
		posts, _ := handler.Repo.GetPostsByPrivacy("rivate")

		userJSON, err := json.Marshal(posts)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts/privacy/rivate"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(userJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the userByIdHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	}
}

func TestPostsByPrivacyHandlerInValidMethodeExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewPostsByPrivacyHandler(R)
		posts, _ := handler.Repo.GetPostsByPrivacy("public")

		userJSON, err := json.Marshal(posts)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts/privacy/public"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(userJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the userByIdHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
	}
}
