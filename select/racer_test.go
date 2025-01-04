package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare get times and return faster", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// nice way to instruct to close resources magically at
		// the end of the function for readability purposes
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Race(slowUrl, fastUrl)

		if err != nil {
			t.Errorf("expected nil error, got %q", err)
		}

		if want != got {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("return error if does not respond in specified time", func(t *testing.T) {
		delayedServer := makeDelayedServer(20 * time.Millisecond)

		defer delayedServer.Close()

		url := delayedServer.URL

		_, err := ConfigurableRacer(url, url, 10 * time.Millisecond)

		if err == nil {
			t.Errorf("expected err, got nil")
		}
	})
}

// httptest.NewServer magically finds ope port and creates server on that
// much amazed
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.Write([]byte{})
	}))
}
