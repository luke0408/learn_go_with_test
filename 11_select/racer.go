package main

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(b string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(b)
		close(ch)
	}()
	return ch
}
