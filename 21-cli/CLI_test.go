package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Like wins\n")
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	assertPlayerWin(t, playerStore, "Like")
}

func assertPlayerWin(t *testing.T, playerStore *StubPlayerStore , winner string) {
	t.Helper()

	got := playerStore.winCalls[0]
	want := "Like"

	if got != want {
		t.Errorf("did not record currect winner, got %s want %s", got, want)
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
