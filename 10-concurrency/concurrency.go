package _0_concurrency

import "time"

type WebsiteChecker func(string) bool

// 返回一个 map，由每个 url 检查后得到的 bool 值组成
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}

/*NOTE
`fatal error: concurrent map writes`

race condition（竞争条件）,多个 goroutine 同时写入 map

使用 go test -race 运行测试，会出现 `WARNING: DATA RACE`
Write at 0x00c00011e2d0 by goroutine 9:
Previous write at 0x00c00011e2d0 by goroutine 8:
Read at 0x00c00008e089 by goroutine 7:
Previous write at 0x00c00008e089 by goroutine 10:

同时操作一个内存，资源竞争，死锁
*/