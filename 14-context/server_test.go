package _4_context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}


func TestServer(t *testing.T) {
	data := "hello Jack"
	server := Server(&StubStore{})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	server.ServerHTTP(recorder, request)

	if recorder.Body.String() != data {
		t.Errorf("got %s want %s", recorder.Body.String(), data)
	}
}