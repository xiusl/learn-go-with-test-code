package _0_concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url[:4] != "http" {
		return false
	}
	return true
}


func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://ins.sleen.top/",
		"waat://abcdefg.mnm/",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	expectedResults := map[string]bool {
		"https://google.com": true,
		"https://ins.sleen.top/": true,
		"waat://abcdefg.mnm/": false,
	}

	if !reflect.DeepEqual(actualResults, expectedResults) {
		t.Errorf("got %v want %v", actualResults, expectedResults)
	}
}
