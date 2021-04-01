package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path,"/players/")

	if player == "Like" {
		_, _ = fmt.Fprintf(w, "20")
	}

	if player == "Jack"{
		_, _ = fmt.Fprintf(w, "30")
	}
}

