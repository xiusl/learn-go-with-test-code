package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Score":20},
			{"Name": "Tom", "Score":11}]`)
		defer closeFile()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{Name: "like", Score:20},
			{Name: "Tom", Score:11},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})


	t.Run("Get player score", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Score":20},
			{"Name": "Tom", "Score":11}]`)
		defer closeFile()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("like")
		want := 20

		assertScore(t, got, want)
	})

	t.Run("Store wins for existing player", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Score":20},
			{"Name": "Tom", "Score":11}]`)
		defer closeFile()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("like")

		got := store.GetPlayerScore("like")
		want := 21

		assertScore(t, got, want)
	})

	t.Run("Store wins for not existing player", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
			{"Name": "like", "Score":20},
			{"Name": "Tom", "Score":11}]`)
		defer closeFile()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("rose")

		got := store.GetPlayerScore("rose")
		want := 1

		assertScore(t, got, want)
	})
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}


func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Errorf("could not create tmp file, %v", err)
	}

	_, _ = tmpFile.Write([]byte(initialData))

	removeFile := func() {
		_, _ = os.ReadFile(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("did not expect an error but got one, %v", err)
	}
}

func TestTape_Write(t *testing.T) {
	file, closeFile := createTempFile(t, "abcdefg")
	defer closeFile()

	tape := &Tape{file}

	_, _ = tape.Write([]byte("123"))

	_, _ = file.Seek(0, 0)
	newFileCont, _ := ioutil.ReadAll(file)

	got := string(newFileCont)

	want := "123"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

/*
两种方案
1. 为每个测试创建一个临时文件。*os.File 实现 ReadWriteSeeker。
	好处是它变得更像集成测试，我们真的是从文件系统中读取和写入，所以我们对此更有信心。
	缺点是我们更喜欢单元测试，因为它们更快而且通常更简单。我们还需要做更多关于创建临时文件的工作，然后确保在测试之后删除它们。
2. 使用第三方库。github.com/mattetti 已经编写了一个 filebuffer 库，它实现了我们需要的接口，并且不触及文件系统。
*/