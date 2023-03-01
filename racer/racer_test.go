package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}

func TestRacer(t *testing.T) {
	t.Run("returns the URL of the faster server", func(t *testing.T) {
		slowServer := makeDelayedServer(10 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		timeout := 25 * time.Millisecond
		got, _ := Racer(slowURL, fastURL, timeout)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns an error if both servers do not respond within the timeout",
		func(t *testing.T) {
			timeout := 10 * time.Millisecond
			serverA := makeDelayedServer(20 * time.Millisecond)
			serverB := makeDelayedServer(25 * time.Millisecond)
			defer serverA.Close()
			defer serverB.Close()

			_, err := Racer(serverA.URL, serverB.URL, timeout)

			if err == nil {
				t.Error("expected an error but didn't get one")
			}
		})
}
