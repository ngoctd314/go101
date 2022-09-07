package main

import (
	"fmt"
	"log"
)

var a = func() any {
	fmt.Println("Run first")
	return nil
}()

func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println("Run main")
	log.Panic("abc")
}
