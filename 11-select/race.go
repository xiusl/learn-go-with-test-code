package _1_select

import (
	"fmt"
	"net/http"
	"time"
)

/*
	比较两个 URL 的响应速度，并返回最先响应的 URL，
	如果超过 10s 都没有响应，返回一个 error
*/

const tenSecondTimeout = 10 * time.Second

func WebsiteRace(aURL, bURL string) string {

	aDuration := measureResponseTime(aURL)
	bDuration := measureResponseTime(bURL)

	if aDuration < bDuration {
		return aURL
	}

	return bURL
}

func WebsiteRaceV2(aURL, bURL string) (string, error) {
	return ConfigurableWebsiteRaceV2(aURL, bURL, tenSecondTimeout)
}

func ConfigurableWebsiteRaceV2(aURL, bURL string, timeout time.Duration) (string, error) {
	select {
	case <- ping(aURL):
		return aURL, nil
	case <- ping(bURL):
		return bURL, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timeout waiting for %s and %s", aURL, bURL)
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
