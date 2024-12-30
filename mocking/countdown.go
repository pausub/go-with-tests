package main

import (
	"fmt"
	"io"
	"os"
)

func CountDown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

func init() {
	CountDown(os.Stdout)
}
