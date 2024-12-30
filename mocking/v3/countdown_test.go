package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {

	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		CountDown(buffer, &SpyCountDownOperations{})

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
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spy := &SpyCountDownOperations{}

		CountDown(spy, spy)

		wanted := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(wanted, spy.Calls) {
			t.Errorf("wanted calls: %q, got: %q", wanted, spy.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := time.Second * 5

	spyTime := &SpyTime{}
	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Should have slept for %v, but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
