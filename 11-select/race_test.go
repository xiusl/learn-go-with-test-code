package _1_select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRace(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := WebsiteRace(slowURL, fastURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}