package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "20")
}

