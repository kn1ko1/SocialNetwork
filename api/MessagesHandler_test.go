package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
	"socialnetwork/repo"
)

func TestMessagesHandlerValidMessageExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {
		// Create a new instance of MessagesHandler with the mock repository
		r := repo.NewDummyRepository()
		handler := NewMessagesHandler(r)

		// Create a sample post to send in the request body
		message := models.GenerateValidMessage()

		messageJSON, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, "/api/messages", bytes.NewBuffer(messageJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the MessagesHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
		// Add additional assertions as needed for your specific use case
	}
}
