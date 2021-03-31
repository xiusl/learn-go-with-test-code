package _3_sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrency", func(t *testing.T) {
		wantCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}

/*NOTE
sync.Mutex，互斥锁，当一个 goroutine 获得 Mutex 后，其它 goroutine 必须等待这个 goroutine 释放这个 Mutex

使用 `go vet` 命令进行语法检查
> ./sync_test.go:15:20: call of assertCounter copies lock value: github.com/xiusl/go-learn/13-sync.Counter contains sync.Mutex
> ./sync_test.go:33:20: call of assertCounter copies lock value: github.com/xiusl/go-learn/13-sync.Counter contains sync.Mutex

需要注意：A Mutex must not be copied after first use. （Mutex 在第一被使用后是不能被拷贝的）
func assertCounter(t testing.TB, got Counter, want int) 函数在使用 got Counter 是值传递，会发生拷贝

*/

/*NOTE
sync.WaitGroup，内部拥有一个计数器，可以添加 n 个并发任务，每个任务完成，计数器减 1，
在另一个 goroutine 等待计数器为 0 时，所有任务完成
保证多个任务的同步，并发中完成指定数量的任务
*/