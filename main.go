package main

import (
	"fmt"
	_ "net/http/pprof"
	"time"
)

func main() {
	fmt.Println(1)
	<-afterDuration(time.Second)
	fmt.Println(2)
	<-afterDuration(time.Second)
	fmt.Println(3)
}

func afterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)

	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()

	return c
}
