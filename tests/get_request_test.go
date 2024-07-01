package web_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGEThello(t *testing.T) {
	t.Run("returns hello", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		GreetingServer(response, request)

		got := response.Body.String()
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}