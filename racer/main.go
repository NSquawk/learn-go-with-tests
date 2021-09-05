package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	// select allows us to wait on multiple channels
	select {
		case <-ping(a)
			return a
		case <-ping(b)
			return b
	}
}

func measureRequestTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

func ping(url string) chan struct{}{
	ch := make(chan struct{})
	go func(){
		http.Get(url)
		//don't care about channel data... just want to signal that Get is done
		close(ch)
	}()
	return ch
}

