package main

import "fmt"

func main() {

	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fn()
}

func fn() {
	panic("acb")
}
