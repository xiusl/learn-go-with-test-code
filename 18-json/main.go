package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", http.HandlerFunc(server.ServerHTTP)); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}