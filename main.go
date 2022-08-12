package main

import "github.com/ngoctd314/go101/go-code/golog"

func main() {
	golog.Info(golog.Sys, "hello world")
	golog.Warn(golog.Sys, "hello world")
	golog.Error(golog.Sys, "hello world")
}
