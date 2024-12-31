package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	counter++
}

// to detect possible race condition, run with go run -race race.go
func main() {

	go increment()
	go increment()	

	time.Sleep(1 * time.Second)
	fmt.Printf("Counter value: %d\n", counter)
}
