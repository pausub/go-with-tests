package main

import (
	"net/http"
	"time"
)

func Race(url1, url2 string) (result string) {
	duration1 := measureResponseTime(url1)
	duration2 := measureResponseTime(url2)

	if duration1 > duration2 {
		return url2
	} else {
		return url1
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
