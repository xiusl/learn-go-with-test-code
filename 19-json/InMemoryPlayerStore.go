package main

import (
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

func (s *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range s.store{
		league = append(league, Player{name, wins})
	}
	return league
}
