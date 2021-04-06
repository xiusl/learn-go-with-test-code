package main

import (
	"fmt"
	poker "github.com/xiusl/go-learn/21-cli"
	"log"
	"os"
)

const dbFileName = "cli.game.json"

func main() {

	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		 log.Fatal(err)
	}

	defer closeFunc()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	poker.NewCLI(store, os.Stdin).PlayPoker()
}
