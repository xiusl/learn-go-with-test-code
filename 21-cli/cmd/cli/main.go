package main

import (
	poker "github.com/xiusl/go-learn/21-cli"
	"log"
	"os"
)

const dbFileName = "cli.game.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening file %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}
