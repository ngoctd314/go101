package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("profile.pb.gz")
	if err != nil {
		log.Fatal(err)
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	n, m, p := 100, 10000, 10000

	s := f1(n, m, p)
	fmt.Println(s)
}

func f1(n, m, p int) int {
	s := 0
	for i := 0; i < n; i++ {
		s += f2(m, p)
	}
	return s
}

func f2(m, p int) int {
	s := 0
	for i := 0; i < m; i++ {
		s += f3(p)
	}
	return s
}

func f3(p int) int {
	s := 0
	for i := 0; i < p; i++ {
		s += i
	}
	return s
}
