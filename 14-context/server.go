package _4_context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

// Server return a handler for calling store
// 返回一个 store 的处理器
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, store.Fetch())
	}
}