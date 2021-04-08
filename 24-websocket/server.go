package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type Player struct {
	Name string
	Wins int
}

// HTTP Server
type PlayerServer struct {
	store    PlayerStore
	http.Handler
	template *template.Template
	game     Game
}

const JsonContentType = "application/json"
const htmlTemplatePath = "game.html"


func NewPlayerServer(store PlayerStore, game Game) *PlayerServer {
	p := new(PlayerServer)

	tmpl, _ := template.ParseFiles(htmlTemplatePath)

	p.template = tmpl
	p.store = store
	p.game = game

	route := http.NewServeMux()
	route.Handle("/league", http.HandlerFunc(p.leagueHandler))
	route.Handle("/players/", http.HandlerFunc(p.playerHandler))
	route.Handle("/game", http.HandlerFunc(p.gameHandler))
	route.Handle("/ws", http.HandlerFunc(p.wsHandler))
	p.Handler = route

	return p
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func (p *PlayerServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)

	_, numberOfPlayersMsg, _ := conn.ReadMessage()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.game.Start(numberOfPlayers, ioutil.Discard)

	_, winner, _ := conn.ReadMessage()
	p.game.Finish(string(winner))
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	_ = p.template.Execute(w, nil)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	_, _ = fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}