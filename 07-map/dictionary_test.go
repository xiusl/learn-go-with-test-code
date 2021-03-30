package _7_map

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "one one one"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "one one one"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("abc")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	dict.Add("test", "one one one")

	want := "one one one"
	got, err := dict.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
