package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
	}
	_, _ = fmt.Fprint(out, finalWord)
}


func main() {
	Countdown(os.Stdout)
}

