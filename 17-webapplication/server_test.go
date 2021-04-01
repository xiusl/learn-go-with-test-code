package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]string
}

func (s *StubPlayerStore)GetPlayerScore(name string) string {
	score := s.scores[name]
	return score
}

func TestPlayerServer(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]string {
			"Like": "20",
			"Jack": "30",
		},
	}
	playerServer := &PlayerServer{store}

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

	t.Run("return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("abc")
		recorder := httptest.NewRecorder()

		playerServer.ServerHTTP(recorder, request)

		assertResponseStatus(t, recorder.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]string{},
	}
	server := &PlayerServer{store}

	t.Run("return accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Like", nil)
		recorder := httptest.NewRecorder()

		server.ServerHTTP(recorder, request)

		assertResponseStatus(t, recorder.Code, http.StatusAccepted)
	})
}


func newGetScoreRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	return request
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status: got %d want %d", got, want)
	}
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