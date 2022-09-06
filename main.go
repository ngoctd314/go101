package main

import (
	"fmt"
	"runtime"
	"strings"
)

// In this case, the `t := T{}` can not measured by this method.
func memUsage(m1, m2 *runtime.MemStats) {
	fmt.Println("Alloc:", m2.Alloc-m1.Alloc,
		"TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
		"HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc,
		"Mallocs: ", m2.Mallocs-m1.Mallocs,
		"HeapInuse: ", m2.HeapInuse-m1.HeapInuse)
}

var a string

func main() {
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	var arr = make([]byte, 10000, 10000)
	a = string(arr)
	fn("abcdefet")

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	memUsage(&memStat, &mem)
}

func fn(s string) {
	// a = s[:5]

	// a = string([]byte(s[:5]))

	var b strings.Builder
	b.Grow(5)
	b.WriteString(s[:5])
	a = b.String()
}
