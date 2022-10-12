package main

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

var cronn *cron.Cron

func prune() {
	fmt.Println("RUN once")
	cronn.AddFunc("@every 0h0m1s", func() {
		fmt.Println("RUN every second")
	})
}

var once sync.Once

func newFn() {
	once.Do(prune)

	fmt.Println("create new instance")
}

func main() {
	cronn = cron.New()
	cronn.Start()

	for i := 0; i < 10; i++ {
		go newFn()
	}
	done := make(chan bool)
	<-done
}
