package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	t.Run("return Score of Like", func(t *testing.T) {
		request := newGetScoreRequest("Like")
		recorder := httptest.NewRecorder()

		PlayerServer(recorder, request)


		assertResponseBody(t, recorder.Body.String(), "20")
	})

	t.Run("return Score of Jack", func(t *testing.T) {
		request := newGetScoreRequest("Jack")
		recorder := httptest.NewRecorder()

		PlayerServer(recorder, request)

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
