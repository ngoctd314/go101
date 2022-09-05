package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		// Simulate a workload
		time.Sleep(time.Second * 5)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	n := time.Now()
	rand.Seed(time.Now().UnixNano())

	a, b := fn(longTimeRequest(), longTimeRequest())

	fmt.Println(sumSquares(a, b))
	fmt.Println(time.Since(n).Seconds())
}

func fn(ch1 <-chan int32, ch2 <-chan int32) (int32, int32) {
	return <-ch1, <-ch2
}
