package _0_concurrency

type WebsiteChecker func(string) bool

// 返回一个 map，由每个 url 检查后得到的 bool 值组成
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

