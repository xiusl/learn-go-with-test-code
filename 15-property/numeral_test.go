package _5_property

import "testing"

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