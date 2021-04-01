package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path,"/players/")

	_, _ = fmt.Fprintf(w, getPlayerScore(player))
}

func getPlayerScore(name string) string {
	switch name {
	case "Like":
		return "20"
	case "Jack":
		return "30"
	default:
		return "0"
	}
}

