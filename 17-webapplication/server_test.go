package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	t.Run("return Score of Like", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Like", nil)
		recorder := httptest.NewRecorder()

		PlayerServer(recorder, request)

		got := recorder.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return Score of Jack", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Jack", nil)
		recorder := httptest.NewRecorder()

		PlayerServer(recorder, request)

		got := recorder.Body.String()
		want := "30"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
