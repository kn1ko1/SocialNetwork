package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/Server/models"
	"strconv"
	"testing"
)

func TestMessageByIdHandlerValidIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewMessageByIdHandler(R)
		messageId := RandomNumberStr

		URL := "/api/messages/" + messageId

		// Create a new HTTP request
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

func TestMessageByIdHandlerValidMessageExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewMessageByIdHandler(R)
		message := models.GenerateValidMessage()
		messageId := strconv.Itoa(message.MessageId)

		messageJSON, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/comments/" + messageId

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(messageJSON))
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

func TestMessageByIdHandlerInValidMessageExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewMessageByIdHandler(R)
		message := models.GenerateInvalidMessage()
		messageId := strconv.Itoa(message.MessageId)

		messageJSON, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/comments/" + messageId

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(messageJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the Handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	}
}

func TestMessageByIdHandlerValidIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {
		handler := NewMessageByIdHandler(R)
		messageId := RandomNumberStr

		URL := "/api/messages/" + messageId

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodDelete, URL, nil)
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

func TestMessageByIdHandlerInValidMethodExpectPass_Post1(t *testing.T) {
	handler := NewMessageByIdHandler(R)
	messageId := RandomNumberStr

	URL := "/api/messages/" + messageId

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the Handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}

func TestMessageByIdHandlerInValidMethodExpectPass_Post2(t *testing.T) {
	handler := NewMessageByIdHandler(R)
	message := models.GenerateValidMessage()
	messageId := strconv.Itoa(message.MessageId)

	messageJSON, err := json.Marshal(message)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/comments/" + messageId

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(messageJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request using the Handler
	handler.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}
