package poker

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

type GameSpy struct {
	StartCalled  bool
	StartCalledWith int
	BlindAlert []byte

	FinishCalled bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
	_, _ = out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishCalledWith = winner
}

func UserSends(message ...string) io.Reader {
	return strings.NewReader(strings.Join(message, "\n"))
}

type StubPlayerStore struct {
	Scores 		map[string]int
	WinCalls 	[]string
	League 		League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Errorf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q wantn %q", store.WinCalls[0], winner)
	}
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at, %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{At: at, Amount: amount})
}
