package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"strconv"
	"testing"
)

func TestPostByIdHandler_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewPostByIdHandler(R)
		postId := RandomNumberStr

		URL := "/api/posts/" + postId

		req, err := http.NewRequest(http.MethodGet, URL, nil)
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

func TestPostByIdHandlerValidPostExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewPostByIdHandler(R)
		post := models.GenerateValidPost()

		postJSON, err := json.Marshal(post)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts/" + strconv.Itoa(post.PostId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(postJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the commentByIdHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	}
}

func TestPostByIdHandlerInValidPostExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewPostByIdHandler(R)
		post := models.GenerateInvalidPost()

		postJSON, err := json.Marshal(post)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/posts/" + strconv.Itoa(post.PostId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(postJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the commentByIdHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	}
}
