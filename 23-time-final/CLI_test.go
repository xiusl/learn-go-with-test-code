package poker_test

import (
	"bytes"
	poker "github.com/xiusl/go-learn/23-time-final"
	"io"

	"strings"
	"testing"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdin = &bytes.Buffer{}
var dummyStdout = &bytes.Buffer{}

type GameSpy struct {
	StartCalled  bool
	StartCalledWith int

	FinishCalled bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishCalledWith = winner
}

func userSends(message ...string) io.Reader {
	return strings.NewReader(strings.Join(message, "\n"))
}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 player and finish game with 'Like' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Like wins")
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCallWith(t, game, "Like")
	})

	t.Run("start game with 8 player and record 'Like' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Like wins")
		cli := poker.NewCLI(in, dummyStdout, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCallWith(t, game, "Like")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStart(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("8", "Lloyd is a killer")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertGameNotFinish(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func assertGameStartedWith(t *testing.T, game *GameSpy, numberOfPlayersWant int) {
	t.Helper()

	if game.StartCalledWith != numberOfPlayersWant {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWant, game.StartCalledWith)
	}
}

func assertGameNotFinish(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.FinishCalled {
		t.Errorf("game should not have finished")
	}
}

func assertGameNotStart(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertFinishCallWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishCalledWith != winner {
		t.Errorf("expected finish called with %q, but got %q", winner, game.FinishCalledWith)
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if want != got {
		t.Errorf("got %q send to stdout but expected %+v", got, messages)
	}
}

func assertScheduleAt(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()

	if got.Amount != want.Amount {
		t.Errorf("got amount %d want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got schedule time of %v, want %v", got.At, want.At)
	}
}