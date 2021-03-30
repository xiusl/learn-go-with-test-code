package _struct

import "testing"

func checkResult(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestPerimeter(t *testing.T) {

	t.Run("Rect Perimeter", func(t *testing.T) {
		rect := Rectangle{10, 10}
		got := rect.Perimeter()
		want := 40.0

		checkResult(t, got, want)
	})

	t.Run("Circle Perimeter", func(t *testing.T) {
		circle := Circle{5}
		got := circle.Perimeter()
		want := 31.41592653589793

		checkResult(t, got, want)
	})
}


func TestArea(t *testing.T) {
	t.Run("Rect Area", func(t *testing.T) {
		rect := Rectangle{6, 7}
		got := rect.Area()
		want := 42.0

		checkResult(t, got, want)
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		checkResult(t, got, want)
	})
}

/*NOTE
	format %.2f -> float64 类型，保留两位小数
*/