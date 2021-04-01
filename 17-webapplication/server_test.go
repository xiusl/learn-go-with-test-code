package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Like", nil)
	recorder := httptest.NewRecorder()

	PlayerServer(recorder, request)

	got := recorder.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
