package _5_property

import (
	"fmt"
	"testing"
)

var testCases = []struct{
	Arabic int
	Roman string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvert2Roman(t *testing.T) {

	for _, tc := range testCases {
		name := fmt.Sprintf("%d convert to %s", tc.Arabic, tc.Roman)
		t.Run(name, func(t *testing.T) {
			got := ConvertToRoman(tc.Arabic)
			want := tc.Roman

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tc := range testCases {
		name := fmt.Sprintf("%s convert to %d", tc.Roman, tc.Arabic)
		t.Run(name, func(t *testing.T) {
			got := ConvertToArabic(tc.Roman)
			want := tc.Arabic

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

/*
func TestNumeral(t *testing.T) {

	testCases := []struct {
		name string
		arabic int
		want string
	}{
		{
			name: "1 -> I",
			arabic: 1,
			want: "I",
		},
		{
			name: "2 -> II",
			arabic: 2,
			want: "II",
		},
		{
			name: "3 -> III",
			arabic: 3,
			want: "III",
		},
		{
			name: "4 -> IV",
			arabic: 4,
			want: "IV",
		},
		{
			name: "10 -> X",
			arabic: 10,
			want: "X",
		},
		{
			name: "14 -> XIV",
			arabic: 14,
			want: "XIV",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvertToRoman(tc.arabic)
			want := tc.want

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}

*/
/*
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
}*/