package slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	//t.Run("collection of 5 numbers", func(t *testing.T){
	//	numbers := [5]int{1, 2, 3, 4, 5}
	//
	//	got := Sum(numbers)
	//	want := 15
	//
	//	if want != got {
	//		t.Errorf("got %d want %d given, %v", got, want, numbers)
	//	}
	//})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 8})
	want := []int{3, 8}

	// slice can only be compared to nil
	//if got != want {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 8})
	want := []int{2, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

/* NOTE:
	reflect.DeepEqual() 不是类型安全的，下面的代码会通过编译，使用时要注意
	a := "abc"
	b := []int{1, 2}
	if reflect.DeepEqual(a, b) {
		fmt.Println("equal")
	}
*/