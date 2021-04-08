package poker_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	poker "github.com/xiusl/go-learn/24-websocket"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestGETPlayer(t *testing.T) {
	store := poker.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := poker.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("GET /game returns 200", func(t *testing.T) {
		server := poker.NewPlayerServer(&poker.StubPlayerStore{})

		request := newGameRequest()
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, request)

		assertStatus(t, recorder.Code, http.StatusOK)
	})

	t.Run("when we got a message over a websocket it is a winner of a game", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		winner := "Like"
		server := httptest.NewServer(mustMakePlayerServer(t, store))
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("Could not open a ws connection on %s %v", wsURL, err)
		}
		defer func() {
			_ = ws.Close()
		}()

		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		poker.AssertPlayerWin(t, store, winner)
	})

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer func() {
			_ = ws.Close()
		}()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCallWith(t, game, winner)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns the League table as JSON", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Like", 20},
			{"Tom", 11},
		}

		store := &poker.StubPlayerStore{League: wantedLeague}
		server := poker.NewPlayerServer(store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, request)

		got := getLeagueFromResponse(t, recorder.Body)

		assertContentType(t, recorder, poker.JsonContentType)
		assertStatus(t, recorder.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore) *poker.PlayerServer {
	server := poker.NewPlayerServer(store)
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
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

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league poker.League) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Errorf("Unable to parse response from server %s, err: %v", body, err)
	}
	return
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Header().Get("content-type") != want {
		t.Errorf("Response did not have content-type of %s, got %v", want, response.HeaderMap)
	}
}

func assertLeague(t *testing.T, got, want poker.League) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}