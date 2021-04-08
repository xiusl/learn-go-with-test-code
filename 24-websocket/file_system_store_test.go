package poker_test

import (
	poker "github.com/xiusl/go-learn/24-websocket"
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, _ = tmpFile.Write([]byte(initialData))

	removeFile := func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func TestFileSystemStore(t *testing.T) {

	t.Run("League sorted", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Wins":2},
			{"Name": "Tom", "Wins":11}]`)
		defer closeFile()

		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		want := poker.League{
			{Name: "Tom", Wins: 11},
			{Name: "like", Wins: 2},
		}

		assertLeague(t, got, want)
	})

	t.Run("Works with an empty file", func(t *testing.T) {
		database, closeFile := createTempFile(t, "")
		defer closeFile()

		_, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})

	t.Run("League from a reader", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Wins":20},
			{"Name": "Tom", "Wins":11}]`)
		defer closeFile()

		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := []poker.Player{
			{Name: "like", Wins: 20},
			{Name: "Tom", Wins: 11},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("Get player score", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Wins":20},
			{"Name": "Tom", "Wins":11}]`)
		defer closeFile()

		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("like")
		want := 20

		assertScoreEquals(t, got, want)
	})

	t.Run("Store wins for existing player", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Wins":20},
			{"Name": "Tom", "Wins":11}]`)
		defer closeFile()

		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("like")

		got := store.GetPlayerScore("like")
		want := 21

		assertScoreEquals(t, got, want)
	})

	t.Run("Store wins for not existing player", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Wins":20},
			{"Name": "Tom", "Wins":11}]`)
		defer closeFile()

		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("rose")

		got := store.GetPlayerScore("rose")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
