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

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
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
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
