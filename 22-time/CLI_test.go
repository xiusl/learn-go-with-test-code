package poker

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)


func TestCLI(t *testing.T) {
	var dummyBlindAlerter = &SpyBlindAlerter{}
	var dummyPlayerStore = &StubPlayerStore{}
	var dummyStdIn = &bytes.Buffer{}
	// 提示用户输入玩家的数量
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		game := NewGame(dummyBlindAlerter, dummyPlayerStore)
		cli := NewCLI(dummyStdIn, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompt
		if got != want {
			t.Errorf("go %q want %q", got, want)
		}
	})

	//
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(blindAlerter, dummyPlayerStore)

		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompt
		if got != want {
			t.Errorf("go %q want %q", got, want)
		}

		testCases := []ScheduleAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, tc := range testCases {
			t.Run(fmt.Sprint(tc), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alter %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, tc)
			})
		}
	})
}



func assertScheduledAlert(t *testing.T, got, want ScheduleAlert) {
	t.Helper()

	if got.Amount != want.Amount {
		t.Errorf("got amount %d want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got schedule time of %v, want %v", got.At, want.At)
	}
}

/*
t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Like wins")
		playerStore := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		testCases := []scheduleAlert{
			{0 * time.Minute, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, tc := range testCases {
			t.Run(fmt.Sprintf("%d scheduled for %v", tc.amount, tc.at), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alter %d was not scheduled %v", i, blindAlerter.alerts)
				}

				alter := blindAlerter.alerts[i]

				assertScheduledAlert(t, alter, tc)
			})
		}
	})
*/
/*
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
*/
