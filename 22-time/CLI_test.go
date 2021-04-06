package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Like wins")
		playerStore := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Errorf("expected a blind alter to be scheduled")
		}
	})
}

func assertPlayerWin(t *testing.T, playerStore *StubPlayerStore , winner string) {
	t.Helper()

	got := playerStore.winCalls[0]
	if got != winner {
		t.Errorf("did not record currect winner, got %s want %s", got, winner)
	}
}

/*
func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Errorf("expected a win call but didn't get any")
	}
}
*/
