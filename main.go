package main

import (
	"strconv"

	"github.com/ngoctd314/go101/go-code/golog"
)

func main() {
	fn := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}
	k := fn()
	golog.Info(golog.Sys, strconv.Itoa(k()))
	golog.Info(golog.Sys, strconv.Itoa(k()))
	golog.Info(golog.Sys, strconv.Itoa(k()))
	golog.Info(golog.Sys, strconv.Itoa(k()))

	var a = 1
	switch a {
	case 1:
		golog.Info(golog.Sys, strconv.Itoa(a))
		fallthrough
	case 2:
		golog.Info(golog.Sys, strconv.Itoa(a))
	}
}
