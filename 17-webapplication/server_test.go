package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	playerServer := &PlayerServer{}

	t.Run("return Score of Like", func(t *testing.T) {
		request := newGetScoreRequest("Like")
		recorder := httptest.NewRecorder()

		playerServer.ServerHTTP(recorder, request)


		assertResponseBody(t, recorder.Body.String(), "20")
	})

	t.Run("return Score of Jack", func(t *testing.T) {
		request := newGetScoreRequest("Jack")
		recorder := httptest.NewRecorder()

		playerServer.ServerHTTP(recorder, request)

		assertResponseBody(t, recorder.Body.String(), "30")
	})
}

func newGetScoreRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	return request
}


func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
--- FAIL: TestPlayerServer (0.00s)
    --- FAIL: TestPlayerServer/return_Score_of_Like (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
        panic: runtime error: invalid memory address or nil pointer dereference

playerServer := &PlayerServer{} not PlayerStore ??
*/