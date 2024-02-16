package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroupUserByGroupIdAndUserIdHandlerValidIdsExpectPass_Delete(t *testing.T) {
	for i := 0; i < 10; i++ {

		handler := NewGroupUserByGroupIdAndUserIdHandler(R)
		groupId := RandomNumberStr
		userId := RandomNumberStr

		URL := "/api/groups/" + groupId + "/eventUsers/users/" + userId

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

func TestGroupUserByGroupIdAndUserIdHandlerInValidMethodPostExpectPass_Post(t *testing.T) {
	handler := NewGroupUserByGroupIdAndUserIdHandler(R)
	groupId := RandomNumberStr
	userId := RandomNumberStr

	URL := "/api/groups/" + groupId + "/eventUsers/users/" + userId

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

func TestGroupUserByGroupIdAndUserIdHandlerInValidMethodPutExpectPass_Put(t *testing.T) {
	handler := NewGroupUserByGroupIdAndUserIdHandler(R)
	groupId := RandomNumberStr
	userId := RandomNumberStr

	URL := "/api/groups/" + groupId + "/eventUsers/users/" + userId

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
