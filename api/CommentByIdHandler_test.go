package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestUsersHandlerValidUserExpectPass_Post
func TestCommentByIdHandlerValidIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewCommentByIdHandler(R)
		comment, _ := handler.Repo.GetCommentById(rand.Intn(101))

		commentJSON, err := json.Marshal(comment)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/comments/" + fmt.Sprint(comment.CommentId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(commentJSON))
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
		comment, _ := handler.Repo.GetCommentById(rand.Intn(101))

		commentJSON, err := json.Marshal(comment)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/comments/" + fmt.Sprint(comment.CommentId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(commentJSON))
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

func TestCommentByIdHandlerValidIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewCommentByIdHandler(R)
		comment, _ := handler.Repo.GetCommentById(rand.Intn(101))
		err := handler.Repo.DeleteCommentById(comment.CommentId)

		if err != nil {
			t.Fatal(err)
		}

		commentJSON, err := json.Marshal(comment)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/comments/" + fmt.Sprint(comment.CommentId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodDelete, URL, bytes.NewBuffer(commentJSON))
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
