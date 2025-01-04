package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsDuration = 10 * time.Second

func Race(url1, url2 string) (result string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondsDuration)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (result string, err error) {

	// select allows to wait on multipe channels. First one
	// to send value - wins and its code block gets executed
	select {

	// reader unblocks when receives data or when channel
	// is closed
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil

	// wait at most 10 seconds, then return error.
	case <-time.After(timeout):
		return "", fmt.Errorf("time out after waiting for %d waiting for %q or %q", timeout, url1, url2)
	}
}

func ping(url string) chan struct{} {
	// always create channel with make, othervice it will be
	// nil and wont accept messages
	// empty struct is cheaper chan boolean channel
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// close channel after get returns
		close(ch)
	}()
	// return channel
	return ch
}
