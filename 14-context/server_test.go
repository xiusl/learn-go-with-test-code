package _4_context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("Tells store cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello Jack"
		store := &StubStore{response: data}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(10 * time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, req)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		data := "Hello Jack"
		store := &StubStore{response: data}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, req)

		if recorder.Body.String() != data {
			t.Errorf("got %s want %s", recorder.Body.String(), data)
		}

		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})
}

/*NOTE
context 上下文，主要用来在 goroutine 间传递信息，如：取消信号、超时时间、截止时间等

> The context package provides functions to derive new Context values from existing ones.
> These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.

context 包提供了一个函数，可以基于一个存在的上下文（req.Context）产生一个新的上下文（cancellingCtx），
派生出的上下文会形成一个 `树`，当一个上下文被取消时，从这个上派生的上下文都会被取消
*/

/*Version 1
func TestServer(t *testing.T) {
	data := "hello Jack"
	server := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	server.ServeHTTP(recorder, request)

	if recorder.Body.String() != data {
		t.Errorf("got %s want %s", recorder.Body.String(), data)
	}
}
*/