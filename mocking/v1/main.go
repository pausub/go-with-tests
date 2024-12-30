package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countFrom = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

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
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
