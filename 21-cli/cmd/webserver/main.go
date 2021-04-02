package main

import (
	"log"
	"net/http"
	"os"

	"github.com/xiusl/go-learn/21-cli"
)

const dbFileName = "com.xiusl.like"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening file %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server.Handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
