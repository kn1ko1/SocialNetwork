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

func TestEventsByGroupIdHandlerValidGroupIdExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventsByGroupIdHandler(R)
		event, _ := handler.Repo.GetEventById(rand.Intn(101))
		groupId := event.GroupId

		eventJSON, err := json.Marshal(groupId)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

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
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventsByGroupIdHandlerInValidMethodExpectPass_Get(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventsByGroupIdHandler(R)
		event, _ := handler.Repo.GetEventById(rand.Intn(101))
		groupId := event.GroupId

		eventJSON, err := json.Marshal(groupId)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

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
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventsByGroupIdHandlerValidGroupIdExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventsByGroupIdHandler(R)
		event, _ := handler.Repo.GetEventById(rand.Intn(101))
		groupId := event.GroupId

		eventJSON, err := json.Marshal(groupId)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

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
		// Add additional assertions as needed for your specific use case
	}
}

func TestEventsByGroupIdHandlerInValidMethodExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewEventsByGroupIdHandler(R)
		event, _ := handler.Repo.GetEventById(rand.Intn(101))
		groupId := event.GroupId

		eventJSON, err := json.Marshal(groupId)
		if err != nil {
			t.Fatal(err)
		}

		URL := "/api/groups/" + fmt.Sprint(groupId) + "/events"

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
		// Add additional assertions as needed for your specific use case
	}
}
