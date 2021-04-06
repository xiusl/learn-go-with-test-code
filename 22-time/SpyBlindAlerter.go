package poker

import "time"

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type SpyBlindAlerter struct {
	alerts []struct{
		scheduledAt time.Duration
		amount int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{scheduledAt: duration, amount: amount})
}