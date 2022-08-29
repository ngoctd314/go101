package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Timer
	t = *time.NewTimer(time.Second * 5)

	select {
	case v := <-t.C:
		fmt.Println(v)
	}
}
