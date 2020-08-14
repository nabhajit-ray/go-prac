package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	fmt.Println(url)
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
