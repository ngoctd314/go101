package main

import (
	"fmt"
	_ "net/http/pprof"
)

func main() {
	s := []int{1, 2, 3}
	s1 := s[:2:2]
	s2 := append(s1, 10)
	fmt.Println(s, s1, s2)
}
