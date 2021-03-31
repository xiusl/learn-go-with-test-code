package _1_select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRace(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := WebsiteRace(slowURL, fastURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func TestWebsiteRaceV2(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := WebsiteRaceV2(slowURL, fastURL)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})


	t.Run("Timeout", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := WebsiteRaceV2(serverA.URL, serverB.URL)

		if err == nil {
			t.Errorf("excepted an error but did not get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
