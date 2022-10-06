package main

import (
	"fmt"
	_ "net/http/pprof"
	"sync"
)

func main() {
	// The capacity must be one
	mutex := make(chan struct{}, 1)

	counter := 0
	increase := func() {
		// lock through send
		mutex <- struct{}{}
		counter++
		// unlock through receive
		<-mutex
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	increase1000 := func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			increase()
		}
	}

	go increase1000()
	go increase1000()
	wg.Wait()
	fmt.Println(counter)
}
