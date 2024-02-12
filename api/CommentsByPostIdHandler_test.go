package api

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCommentsByPostIdHandlerValidPostIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		// Create a new instance of commentByIdHandler with the mock repository
		handler := NewCommentsByPostIdHandler(R)
		postId := rand.Intn(101)
		postIdStr := strconv.Itoa(postId)

		URL := "/api/posts/" + postIdStr + "/comments"

		// Create a new HTTP request with the encoded JSON as the request body
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
		postId := rand.Intn(101)
		postIdStr := strconv.Itoa(postId)

		URL := "/api/posts/" + postIdStr + "/comments"

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
}
