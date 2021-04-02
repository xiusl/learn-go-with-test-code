package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name string
	Score int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

func (s *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path,"/players/")
	switch r.Method {
	case http.MethodPost:
		s.ProcessWin(w, player)
	case http.MethodGet:
		s.ShowScore(w, player)
	}
}

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(s.getLeagueTable())
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
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

func (s *PlayerServer) getLeagueTable() []Player{
	return s.store.GetLeague()
}