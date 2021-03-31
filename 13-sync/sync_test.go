package _3_sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrency", func(t *testing.T) {
		wantCount := 1000
		counter := Counter{}

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

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}

/*NOTE
sync.WaitGroup，内部拥有一个计数器，可以添加 n 个并发任务，每个任务完成，计数器减 1，
在另一个 goroutine 等待计数器为 0 时，所有任务完成
保证多个任务的同步，并发中完成指定数量的任务
*/