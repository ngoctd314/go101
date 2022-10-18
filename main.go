package main

import (
	"fmt"
	"sync"
)

func main() {
	receiver(sender(10), 10)
}

func sender(max int) <-chan int {
	c := make(chan int, max)
	go func() {
		for i := 0; i < max; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receiver(c <-chan int, max int) {
	wg := sync.WaitGroup{}
	wg.Add(max)
	for i := 0; i < max; i++ {
		// multiple receivers
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Println(v)
			}
		}()
	}
	wg.Wait()
}
