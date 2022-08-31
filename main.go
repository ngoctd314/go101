package main

import (
	"github.com/ngoctd314/go101/go-code/golog"
)

// Person ...
type Person struct {
	Name []byte
}

func main() {
	golog.Errorw(golog.Server, "Hello world", "TDN", "XYZ", "MNP", "P")
	golog.Infow(golog.Test, "Hello world", "abc", "C")
}
