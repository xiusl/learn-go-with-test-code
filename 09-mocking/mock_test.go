package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

/*NOTE
反引号语法是创建 string 的另一种方式，但是允许你放置东西例如放到新的一行
*/