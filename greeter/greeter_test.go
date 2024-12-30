package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	// Buffer implements Writer interface
	// because it has a method Write(..)
	buffer := bytes.Buffer{}

	Greet(&buffer, "Pliupt")

	got := buffer.String()
	want := "Hello, Pliupt"

	if got != want {
		t.Errorf("got %q, want: %q", got, want)
	}
}
