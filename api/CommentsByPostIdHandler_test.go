package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommentsByPostIdHandler_Get(t *testing.T) {

	// Create a new instance of commentByIdHandler with the mock repository
	handler := NewCommentsByPostIdHandler(R)
	comment, _ := handler.Repo.GetAllComments()

	postId := comment[0].PostId

	commentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/posts/" + fmt.Sprint(postId) + "comments"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(commentJSON))
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

func TestCommentsByPostIdHandler_Delete(t *testing.T) {
	handler := NewCommentsByPostIdHandler(R)
	comment, _ := handler.Repo.GetAllComments()

	postId := comment[0].PostId

	commentJSON, err := json.Marshal(postId)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/posts/" + fmt.Sprint(postId) + "comments"

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodDelete, URL, bytes.NewBuffer(commentJSON))
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
