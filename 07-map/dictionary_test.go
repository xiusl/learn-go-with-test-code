package _7_map

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "one one one"}

	t.Run("known word", func(t *testing.T) {
		got, _ := Search(dictionary, "test")
		want := "one one one"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := Search(dictionary, "abc")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertString(t, err.Error(), want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
