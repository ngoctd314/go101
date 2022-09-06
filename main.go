package main

import (
	"fmt"
	"runtime"
	"strings"
)

var a string = "abcdefghijklmnopqrstuv"

func main() {
	var b strings.Builder
	b.Grow(5000)
	fn(a)
	fmt.Println(len(a))
}

var p = fmt.Println

func memUsage(m1, m2 *runtime.MemStats) {
	p("Alloc:", m2.Alloc-m1.Alloc,
		"TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
		"HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}

func fn(s string) {
	a = string([]byte(s[:5]))
}
