package main

import (
	"fmt"
)

func main() {
	ch := make(chan *int)
	go func() {
		p := 10
		ch <- &p
		fmt.Printf("%p\n", &p)
	}()
	v := <-ch
	fmt.Printf("%p\n", v)

}
