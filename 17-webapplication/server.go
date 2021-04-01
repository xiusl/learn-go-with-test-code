package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path,"/players/")

	_, _ = fmt.Fprintf(w, s.store.GetPlayerScore(player))
}


