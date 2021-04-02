package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "like", "Score":20},
			{"Name": "Tom", "Score":11}]`)

		store := FileSystemStore{database}

		got := store.GetLeague()

		want := []Player{
			{Name: "like", Score:20},
			{Name: "Tom", Score:11},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}