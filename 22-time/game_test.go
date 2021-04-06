package poker

import (
	"fmt"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(blindAlerter, &StubPlayerStore{})

		game.Start(5)

		testCases := []ScheduleAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(testCases, t, blindAlerter)
	})

	t.Run("game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(blindAlerter, &StubPlayerStore{})

		game.Start(7)

		cases := []ScheduleAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &StubPlayerStore{}
	game := NewGame(&SpyBlindAlerter{}, store)
	winner := "Ruth"

	game.Finish(winner)
	AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(testCases []ScheduleAlert, t *testing.T, blindAlerter *SpyBlindAlerter) {
	t.Helper()

	for i, tc := range testCases {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alter %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, tc)
		})
	}
}
