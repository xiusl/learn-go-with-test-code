package main

import (
"fmt"
"net/http"
"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path,"/players/")
	switch r.Method {
	case http.MethodPost:
		s.ProcessWin(w, player)
	case http.MethodGet:
		s.ShowScore(w, player)
	}
}

func (s *PlayerServer) ShowScore(w http.ResponseWriter, player string){

	score := s.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprintf(w, "%d", s.store.GetPlayerScore(player))
}

func (s *PlayerServer) ProcessWin(w http.ResponseWriter, player string){

	s.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
