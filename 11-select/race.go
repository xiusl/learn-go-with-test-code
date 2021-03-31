package _1_select

import (
	"net/http"
	"time"
)

/*
	比较两个 URL 的响应速度，并返回最先响应的 URL，
	如果超过 10s 都没有响应，返回一个 error
*/

func WebsiteRace(aURL, bURL string) string {

	aDuration := measureResponseTime(aURL)
	bDuration := measureResponseTime(bURL)

	if aDuration < bDuration {
		return aURL
	}

	return bURL
}

func WebsiteRaceV2(aURL, bURL string) string {
	select {
	case <- ping(aURL):
		return aURL
	case <- ping(bURL):
		return bURL
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		_, _ =http.Get(url)
		ch <- true
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
