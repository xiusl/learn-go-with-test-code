package _5_property

import "testing"

func TestNumeral(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}