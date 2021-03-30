package _struct

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestArea(t *testing.T) {
	got := Area(6, 7)
	want := 42

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

/*NOTE
	format %.2f -> float64 类型，保留两位小数
*/