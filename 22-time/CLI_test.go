package poker

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Like wins")
		playerStore := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		testCases := []struct{
			expectedScheduleTime time.Duration
			expectedAmount int
		} {
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
			t.Run(fmt.Sprintf("%d scheduled for %v", tc.expectedAmount, tc.expectedScheduleTime), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alter %d was not scheduled %v", i, blindAlerter.alerts)
				}

				alter := blindAlerter.alerts[i]

				amountGot := alter.amount
				if amountGot != tc.expectedAmount {
					t.Errorf("got amount %d want %d", amountGot, tc.expectedAmount)
				}

				gotScheduleTime := alter.scheduledAt
				if gotScheduleTime != tc.expectedScheduleTime {
					 t.Errorf("got schedule time of %v, want %v", gotScheduleTime, tc.expectedScheduleTime)
				}
			})
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
