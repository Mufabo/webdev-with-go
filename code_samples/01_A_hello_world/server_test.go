package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGETHelloWorld(t *testing.T) {
	t.Run("GET", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		getHomeHandleFunc(response, request)

		got := response.Body.String()
		want := "Hello World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
