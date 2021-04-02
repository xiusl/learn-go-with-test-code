package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "com.xiusl.like"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening file %s %v", dbFileName, err)
	}

	store := NewFileSystemStore(db)
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server.Handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

