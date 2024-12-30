package main

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("repeat character 10 times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"
		assertCorrectResult(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 6)
	fmt.Println(result)
}

func assertCorrectResult(t testing.TB, got, want string) {
	if got != want {
		t.Helper()
		t.Errorf("got %q, want %q", got, want)
	}
}
