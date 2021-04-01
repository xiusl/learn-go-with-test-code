package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {

}

func (s *InMemoryPlayerStore)GetPlayerScore(name string) string {
	return "120"
}

func (s *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", http.HandlerFunc(server.ServerHTTP)); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}