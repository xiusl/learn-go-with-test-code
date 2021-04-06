package main

import (
	"github.com/xiusl/go-learn/21-cli"
	"log"
	"net/http"
)

const dbFileName = "com.xiusl.like"

func main() {

	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeFunc()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
