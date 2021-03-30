package _7_map

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "one one one"}

	got := Search(dictionary, "test")
	want := "one one one"
	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
