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

	result := doSum()
	fmt.Println(result)

}

func doSum() int {
	sum := 0
	for i := 0; i < 787766777; i++ {
		sum += i
	}

	return sum
}
