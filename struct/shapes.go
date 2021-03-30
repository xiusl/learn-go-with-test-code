package _struct

import "math"

type Rectangle struct {
	 Width 	float64
	 Height float64
}

func (rect Rectangle)Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle)Area() float64 {
	return rect.Width * rect.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle)Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle)Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}


/* Version 2
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
*/

/* Version 1
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}
 */
