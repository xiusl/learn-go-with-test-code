package _struct

import "testing"

func checkResult(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestPerimeter(t *testing.T) {
	rect := Rectangle{10, 10}
	got := Perimeter(rect)
	want := 40.0

	checkResult(t, got, want)
}


func TestArea(t *testing.T) {
	rect := Rectangle{6, 7}
	got := Area(rect)
	want := 42.0

	checkResult(t, got, want)
}

/*NOTE
	format %.2f -> float64 类型，保留两位小数
*/