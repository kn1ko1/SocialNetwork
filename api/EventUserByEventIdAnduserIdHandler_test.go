package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventUserByEventIdAnduserIdHandler_Delete(t *testing.T) {

	handler := NewEventUserByEventIdAndUserIdHandler(R)
	event, _ := handler.Repo.CreateEvent(*EventExample)
	eventId := event.EventId
	userId := event.UserId

	// err := handler.Repo.DeleteEventUserByEventIdAndUserId(eventId, userId)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	eventJSON, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}

	URL := "/api/events/" + fmt.Sprint(eventId) + "/eventUsers/users/" + fmt.Sprint(userId)

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
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
	}
	// Add additional assertions as needed for your specific use case
}
