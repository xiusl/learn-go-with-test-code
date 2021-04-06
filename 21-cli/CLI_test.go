package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record like win from user input", func(t *testing.T) {
		in := strings.NewReader("Like wins\n")
		playerStore := &StubPlayerStore{}
		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Like")
	})

	t.Run("record jack win from user input", func(t *testing.T) {
		in := strings.NewReader("Jack wins\n")
		playerStore := &StubPlayerStore{}
		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Jack")
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
