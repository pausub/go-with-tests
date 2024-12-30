package main

import (
	"log"
	"net/http"

	//	"os"

	"github.com/pausub/go-with-tests/greeter"
)

// func main() {
// 	greeter.Greet(os.Stdout, "Pupup\n")
// }

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greeter.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
