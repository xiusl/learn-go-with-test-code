package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore)GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestLeague(t *testing.T) {
	store := &StubPlayerStore{}
	server := NewPlayerServer(store)

	t.Run("return 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, request)

		var got []Player

		err := json.NewDecoder(recorder.Body).Decode(&got)

		if err != nil {
			t.Errorf("Unable to parse response from server %s, err: %v", recorder.Body, err)
		}

		assertResponseStatus(t, recorder.Code, http.StatusOK)
	})
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status: got %d want %d", got, want)
	}
}