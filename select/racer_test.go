package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20)
	fastServer := makeDelayedServer(0)

	// nice way to instruct to close resources magically at
	// the end of the function for readability purposes
	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Race(slowUrl, fastUrl)

	if want != got {
		t.Errorf("got: %q, want: %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}

// httptest.NewServer magically finds ope port and creates server on that
// much amazed
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Millisecond)
		w.Write([]byte{})
	}))
}
