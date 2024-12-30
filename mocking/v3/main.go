package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countFrom = 3

type SpyCountDownOperations struct {
	Calls []string
}

func (c *SpyCountDownOperations) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	// methods are ideomatic in go, but function fields
	// allow for dynamic function replacement for mocks with a
	// price of an accidental function replacement with invalid one
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countFrom; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(out, finalWord)
}

func main() {
	// need to create spy as a &pointer because it has pointer receiver method
	// (method, that changes struct field value)
	// automatic dereferencing doesnt happen
	// when type passed as interface
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
