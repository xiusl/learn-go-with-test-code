package _5_property

import "testing"

func TestNumeral(t *testing.T) {

	t.Run("1 -> I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("2 -> II", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}