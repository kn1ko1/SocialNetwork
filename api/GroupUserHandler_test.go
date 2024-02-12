package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"socialnetwork/models"
	"testing"
)

func TestGroupUserHandlerValidGroupUserExpectPass_Post(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewGroupUsersHandler(R)

		groupUser := models.GenerateValidGroupUser()

		groupUserJSON, err := json.Marshal(groupUser)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/api/groupUsers", bytes.NewBuffer(groupUserJSON))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	}
}
