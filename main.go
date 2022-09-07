package main

import (
	"fmt"
	"math/rand"
)

func fn() int {
	arr := []int{1, 2, 3, 4}
	return rand.Intn(len(arr))
}

type person struct {
	friends map[string]int
	age     *int
}

var k = 10

func (p person) u() {
	p.friends["b"] = 2
	*p.age = 20
}

func main() {
	p := person{
		friends: map[string]int{
			"a": 1,
		},
		age: &k,
	}
	p.u()
	fmt.Println(*p.age, k)
}
