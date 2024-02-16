package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"socialnetwork/models"
)

func TestMessagesHandlerValidMessageExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewMessagesHandler(R)
		message := models.GenerateValidMessage()

		messageJSON, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/messages"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(messageJSON))
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
	}
}

func TestMessagesHandlerInValidMessageExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewMessagesHandler(R)
		message := models.GenerateInvalidMessage()

		messageJSON, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/messages"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(messageJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the MessagesHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
		// Add additional assertions as needed for your specific use case
	}
}

func TestMessagesHandlerInValidMethodExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewMessagesHandler(R)
		message := models.GenerateValidMessage()

		messageJSON, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/messages"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(messageJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the MessagesHandler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
		// Add additional assertions as needed for your specific use case
	}
}
