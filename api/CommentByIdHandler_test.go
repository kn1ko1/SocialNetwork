package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"testing"
)

// TestUsersHandlerValidUserExpectPass_Post
func TestCommentByIdHandlerValidIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewCommentByIdHandler(R)
		commentId := RandomNumberStr

		URL := "/api/comments/" + commentId

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodGet, URL, nil)
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

func TestCommentByIdHandlerValidIdExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewCommentByIdHandler(R)
		comment, _ := handler.Repo.UpdateComment(*models.GenerateValidComment())

		commentJSON, err := json.Marshal(comment)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/comments/" + fmt.Sprint(comment.CommentId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(commentJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the commentByIdHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	}
}

func TestCommentByIdHandlerInValidCommentExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewCommentByIdHandler(R)
		comment, _ := handler.Repo.UpdateComment(*models.GenerateInvalidComment())

		commentJSON, err := json.Marshal(comment)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/comments/" + fmt.Sprint(comment.CommentId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(commentJSON))
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

func TestCommentByIdHandlerValidIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewCommentByIdHandler(R)
		commentId := RandomNumberStr

		URL := "/api/comments/" + commentId

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodDelete, URL, nil)
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

func TestCommentByIdHandlerInValidMethodExpectPass(t *testing.T) {
	handler := NewCommentByIdHandler(R)
	commentId := RandomNumberStr

	URL := "/api/comments/" + commentId

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the commentByIdHandler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}
