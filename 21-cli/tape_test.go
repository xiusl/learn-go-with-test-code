package poker

import (
	"io/ioutil"
	"testing"
)

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
