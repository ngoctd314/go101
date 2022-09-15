package main

import (
	"fmt"
	"runtime"
)

func notify(msg *string) {
	fmt.Println("msg: ", *msg)
}

func main() {
	var status string
	defer func() {
		fmt.Println("msg: ", status)
	}()
	a := "abc"
	status = a
}

func printAlloc() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("%d KB\n", mem.Alloc/1000)
}
