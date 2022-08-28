package main

import (
	"fmt"
)

func fn() {
	func() {
		defer fmt.Println("abc")
	}()
	fmt.Println("RUN")

}

func main() {
	fn()
}
