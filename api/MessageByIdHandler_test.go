package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"strconv"
	"testing"
)

func TestMessageByIdHandler_Get(t *testing.T) {
	handler := NewMessageByIdHandler(R)
	message, _ := handler.Repo.GetMessageById(1)

	userJSON, err := json.Marshal(message)
	if err != nil {
		t.Fatal(err)
	}

	randomNumber := rand.Intn(5)
	randomNumberStr := strconv.Itoa(randomNumber)
	URL := "/api/messages/" + randomNumberStr

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(userJSON))
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

func TestMessageByIdHandler_Put(t *testing.T) {
	handler := NewMessageByIdHandler(R)
	notifcation, _ := handler.Repo.UpdateMessage(*models.GenerateValidMessage())
	notifcationJSON, err := json.Marshal(notifcation)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/comments/" + fmt.Sprint(notifcation.MessageId)

	// Create a new HTTP request with the encoded JSON as the request body
	req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(notifcationJSON))
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
