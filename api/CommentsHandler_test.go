package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"socialnetwork/repo"
	"testing"
)

func TestCommentsHandlerValidPostExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		// Create a new instance of PostsHandler with the mock repository
		r := repo.NewDummyRepository()
		handler := NewCommentsHandler(r)

		// Create a sample post to send in the request body
		comment := models.GenerateValidComment()

		commentJSON, err := json.Marshal(comment)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/api/comments", bytes.NewBuffer(commentJSON))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json") // Set Content-Type to application/json

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

// func TestCommentsHandler_Get(t *testing.T) {

// 	handler := NewCommentsHandler(R)
// 	comment, _ := handler.Repo.GetAllComments()

// 	commentJSON, err := json.Marshal(comment)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	URL := "/api/comments"

// 	// Create a new HTTP request with the encoded JSON as the request body
// 	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(commentJSON))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a response recorder to capture the response
// 	recorder := httptest.NewRecorder()

// 	// Serve the HTTP request using the handler
// 	handler.ServeHTTP(recorder, req)

// 	// Check the response status code
// 	if recorder.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
// 	}
// }
