package main

import (
	"log"
	"net/http"
	"sync"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

type InMemoryPlayerStore struct {
	mu sync.Mutex
	store map[string]int
}

func (s *InMemoryPlayerStore)GetPlayerScore(name string) int {
	return 	s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[name]++
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", http.HandlerFunc(server.ServerHTTP)); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}