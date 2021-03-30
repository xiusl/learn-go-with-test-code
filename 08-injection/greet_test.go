package _8_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jack")

	got := buffer.String()
	want := "Hello, Jack"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}