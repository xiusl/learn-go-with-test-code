package _2_reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Jack"
	var got []string

	x := struct{
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want 1", len(got))
	}
}
