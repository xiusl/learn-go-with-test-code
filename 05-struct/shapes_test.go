package _5_struct

import "testing"

func checkResult(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		got := shape.Perimeter()
		checkResult(t, got, want)
	}

	t.Run("Rect Perimeter", func(t *testing.T) {
		rect := Rectangle{10, 10}
		checkPerimeter(t, rect, 40.0)
	})

	t.Run("Circle Perimeter", func(t *testing.T) {
		circle := Circle{5}
		checkPerimeter(t, circle, 31.41592653589793)
	})
}


func TestArea(t *testing.T) {

	testCases:= []struct{
		name string
		shape Shape
		area float64
	}{
		{
			name: "Rect Area",
			shape: Rectangle{6, 7},
			area: 42.0,
		},
		{
			name: "Circle Area",
			shape: Circle{10},
			area: 314.1592653589793,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.area {
				t.Errorf("got %.2f want %.2f", got, tc.area)
			}
		})
	}
}

/*NOTE
	format %.2f -> float64 类型，保留两位小数
*/

/*Version 1
func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		checkResult(t, got, want)
	}

	t.Run("Rect Area", func(t *testing.T) {
		rect := Rectangle{6, 7}
		checkArea(t, rect, 42.0)
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
*/