package api

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGroupUserByGroupIdAndUserIdHandlerValidIdsExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

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
}

func TestGroupUserByGroupIdAndUserIdHandlerInValidMethodPostExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewGroupUserByGroupIdAndUserIdHandler(R)

		groupId := rand.Intn(101)
		groupIdStr := strconv.Itoa(groupId)
		userId := rand.Intn(101)
		userIdStr := strconv.Itoa(userId)
		URL := "/api/groups/" + groupIdStr + "/eventUsers/users/" + userIdStr

		req, err := http.NewRequest(http.MethodPost, URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
	}
}

func TestGroupUserByGroupIdAndUserIdHandlerInValidMethodPutExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewGroupUserByGroupIdAndUserIdHandler(R)

		groupId := rand.Intn(101)
		groupIdStr := strconv.Itoa(groupId)
		userId := rand.Intn(101)
		userIdStr := strconv.Itoa(userId)
		URL := "/api/groups/" + groupIdStr + "/eventUsers/users/" + userIdStr

		req, err := http.NewRequest(http.MethodPut, URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, recorder.Code)
		}
	}
}
