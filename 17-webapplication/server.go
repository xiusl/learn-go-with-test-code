package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		s.ProcessWin(w, r)
	case http.MethodGet:
		s.ShowScore(w, r)
	}
}


func (s *PlayerServer) ShowScore(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path,"/players/")

	score := s.store.GetPlayerScore(player)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprintf(w, s.store.GetPlayerScore(player))
}

func (s *PlayerServer) ProcessWin(w http.ResponseWriter, r *http.Request){
	s.store.RecordWin("Like")
	w.WriteHeader(http.StatusAccepted)
}

