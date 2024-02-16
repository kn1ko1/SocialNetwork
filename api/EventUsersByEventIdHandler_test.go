package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventUsersByEventIdHandlerValidEventIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUsersByEventIdHandler(R)
		event, _ := handler.Repo.GetAllEvents()
		eventId := event[0].EventId

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/events/" + fmt.Sprint(eventId) + "/eventUsers"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(eventJSON))
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

func TestEventUsersByEventIdHandlerInValidEventIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUsersByEventIdHandler(R)
		event, _ := handler.Repo.GetAllEvents()
		eventId := event[0].EventId

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/events/" + fmt.Sprint(eventId) + "/eventUsers"

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
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
	}
}

func TestEventUsersByEventIdHandlerValidEventIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUsersByEventIdHandler(R)
		event, _ := handler.Repo.GetAllEvents()
		eventId := event[0].EventId

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/events/" + fmt.Sprint(eventId) + "/eventUsers"

		// Create a new HTTP request with the encoded JSON as the request body
		req, err := http.NewRequest(http.MethodDelete, URL, bytes.NewBuffer(eventJSON))
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

func TestEventUsersByEventIdHandlerInValidEventIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventUsersByEventIdHandler(R)
		event, _ := handler.Repo.GetAllEvents()
		eventId := event[0].EventId

		eventJSON, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/events/" + fmt.Sprint(eventId) + "/eventUsers"

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
