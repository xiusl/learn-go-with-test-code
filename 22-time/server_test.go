package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const jsonContentType = "application/json"

func TestLeague(t *testing.T) {

	t.Run("return 200 on /league", func(t *testing.T) {
		store := &StubPlayerStore{}
		server := NewPlayerServer(store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, request)

		var got []Player

		err := json.NewDecoder(recorder.Body).Decode(&got)

		if err != nil {
			t.Errorf("Unable to parse response from server %s, err: %v", recorder.Body, err)
		}

		assertResponseStatus(t, recorder.Code, http.StatusOK)
	})

	t.Run("return league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Like", 20},
			{"Tom", 11},
		}

		store := &StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, request)

		got := getLeagueFromResponse(t, recorder.Body)

		assertContentType(t, recorder, jsonContentType)
		assertResponseStatus(t, recorder.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}

func newPostWinRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(http.MethodPost, url, nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	url := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	return request
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Errorf("Unable to parse response from server %s, err: %v", body, err)
	}
	return
}

func assertLeague(t *testing.T, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t *testing.T, resp *httptest.ResponseRecorder, want string) {
	t.Helper()
	if resp.Header().Get("content-type") != want {
		t.Errorf("Response did not have content-type of %s, got %v", want, resp.HeaderMap)
	}
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status: got %d want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
