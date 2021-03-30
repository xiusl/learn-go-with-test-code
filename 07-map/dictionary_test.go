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

	t.Run("New word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add("test", "one one one")

		assertError(t, err, nil)
		assertDefinition(t, dict, "test", "one one one")
	})

	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "one one one"
		dict := Dictionary{word: definition}

		err := dict.Add(word, "two two two")

		assertError(t, err, ErrWorkExists)
		assertDefinition(t, dict, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("got '%s' want '%s'", got, definition)
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
