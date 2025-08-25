package selectch

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares the speeds of two servers, returning the fastest", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatal("got an error but didn't expect one")
		}
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
	t.Run("returns an error if the servers don't respond within 10 s", func(t *testing.T) {
		slowServer := makeDelayedServer(12 * time.Second)
		fastServer := makeDelayedServer(11 * time.Second)
		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)
		if err == nil {
			t.Error("expected an error but didn't got one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
