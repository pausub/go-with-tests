package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	CountDown(buffer, spySleeper)

	got := buffer.String()

	// syntax for multiline strings
	want := `3
2
1
Go!
`
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("Got sleeper calls: %d, want: %d", spySleeper.Calls, 3)
	}

}
