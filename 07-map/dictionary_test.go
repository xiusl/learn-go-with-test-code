package _7_map

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "one one one"}

	got := Search(dictionary, "test")
	want := "one one one"
	if got != want {
		t.Errorf("got '%s' want '%s' given '%v'", got, want, dictionary)
	}
}
