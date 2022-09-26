package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	// set CPU profiling (1)
	f, err := os.Create("profile.pb.gz")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	do()
}

func do() {
	// CPU intensive operation (2)
	test := 0
	for i := 0; i < 1e9; i++ {
		test = i
	}
	fmt.Println(test)
}
