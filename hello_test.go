package main
import "testing"

func TestHello(t *testing.T) {
	got := Hello("Paulius")
	want:= "Hello, Paulius"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}