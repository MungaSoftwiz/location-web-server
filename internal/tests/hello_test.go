package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MungaSoftwiz/location-web-server/internal/handlers"
)

func TestGEThello(t *testing.T) {
	t.Run("returns hello", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		handlers.HelloHandler(response, request)

		got := response.Body.String()
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}