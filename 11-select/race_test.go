package _1_select

import "testing"

func TestWebsiteRace(t *testing.T) {
	slowURL := "https://www.apple.com"
	fastURL := "https://www.baidu.com"

	want := fastURL
	got := WebsiteRace(slowURL, fastURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}