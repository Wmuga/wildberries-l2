package main

import (
	"testing"
	"time"
)

var sig = func(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func TestAll(t *testing.T) {
	// Канал должен закрыться ~ через секунду
	start := time.Now()
	<-ChannelMux(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	dur := time.Since(start)
	if dur < time.Second || dur > 2*time.Second {
		t.Error("Долго работал")
	}
}
