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

func TestEventsByIdHandlerValidIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventByIdHandler(R)
		eventId := RandomNumberStr

		URL := "/api/events/" + eventId

		// Create a new HTTP request
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
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventsByIdHandlerValidIdExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventByIdHandler(R)
		event := models.GenerateValidEvent()

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/events/" + strconv.Itoa(event.EventId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(eventJSON))
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
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventsByIdHandlerInValidEventExpectPass_Put(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventByIdHandler(R)
		event := models.GenerateInvalidEvent()

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/events/" + strconv.Itoa(event.EventId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, recorder.Code)
		}
	}
}

func TestEventsByIdHandlerInValidMethodExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventByIdHandler(R)
		event, _ := handler.Repo.GetEventById(rand.Intn(101))

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/events/" + fmt.Sprint(event.EventId)

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(eventJSON))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to capture the response
		recorder := httptest.NewRecorder()

		// Serve the HTTP request using the handler
		handler.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
	}
}
