package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SapSleeper struct {
	Calls int
}

func (s *SapSleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleep Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleep.Sleep()
		_, _ = fmt.Fprintln(out, i)
	}
	sleep.Sleep()
	_, _ = fmt.Fprint(out, finalWord)
}


type ConfigureSleeper struct {
	duration time.Duration
}

func (c *ConfigureSleeper) Sleep() {
	time.Sleep(c.duration)
}

func main() {
	sleeper := &ConfigureSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}

