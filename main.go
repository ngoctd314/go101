package main

import (
	"fmt"
	_ "net/http/pprof"
)

type Person struct {
	Name *string
}

func main() {
	p := Person{Name: new(string)}
	fn(p)
	fmt.Println(*p.Name)
}

func fn(p Person) {
	*p.Name = "abc"
}
