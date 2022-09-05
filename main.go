package main

import "fmt"

func main() {
	var k int = 0
	if 1 == 1 {
		k := 1
		fmt.Println(k)
	} else {
		k := 2
		fmt.Println(k)
	}
	fmt.Println(k)
}
