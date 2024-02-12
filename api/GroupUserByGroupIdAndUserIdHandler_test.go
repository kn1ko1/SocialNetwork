package api

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGroupUserByGroupIdAndUserIdHandlerEndpointExpectPass_Delete(t *testing.T) {
	handler := NewGroupUserByGroupIdAndUserIdHandler(R)

	groupId := rand.Intn(101)
	groupIdStr := strconv.Itoa(groupId)
	userId := rand.Intn(101)
	userIdStr := strconv.Itoa(userId)
	URL := "/api/groups/" + groupIdStr + "/eventUsers/users/" + userIdStr

	req, err := http.NewRequest(http.MethodDelete, URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}
