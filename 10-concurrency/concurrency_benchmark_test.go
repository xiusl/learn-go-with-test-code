package _0_concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

/*
go test -bench=.

goos: darwin
goarch: amd64
pkg: github.com/xiusl/go-learn/10-concurrency
cpu: Intel(R) Core(TM) i5-8500 CPU @ 3.00GHz
BenchmarkCheckWebsites-6               1        2224623064 ns/op
PASS
ok      github.com/xiusl/go-learn/10-concurrency        2.713s

基准测试使用 100 个网址对 CheckWebsites 进行测试，并使用 WebsiteChecker 伪造实现。
slowStubWebsiteChecker 故意放慢速度，暂停 20 毫秒
执行时间大概 2 秒多
*/

