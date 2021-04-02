package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
	league []Player
}

func (s *StubPlayerStore)GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}


func TestLeague(t *testing.T) {

	t.Run("return 200 on /league", func(t *testing.T) {
		store := &StubPlayerStore{}
		server := NewPlayerServer(store)

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

	t.Run("return league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Like", 20},
			{"Tom", 11},
		}

		store := &StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, request)

		var got []Player

		err := json.NewDecoder(recorder.Body).Decode(&got)

		if err != nil {
			t.Errorf("Unable to parse response from server %s, err: %v", recorder.Body, err)
		}

		assertResponseStatus(t, recorder.Code, http.StatusOK)

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v want %v", got, wantedLeague)
		}
	})
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status: got %d want %d", got, want)
	}
}