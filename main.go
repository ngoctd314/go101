package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	rb := rand.Intn(3) + 1
	time.Sleep(time.Duration(rb) * time.Second)
	select {
	case c <- int32(rb):
	default:
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int32, 1)
	for i := 0; i < 5; i++ {
		go source(c)
	}
	rnd := <-c
	fmt.Println(rnd)
}
