package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	cron := cron.New()

	cron.Start()
	time.Sleep(time.Second * 4)
	cron.AddFunc("@every 0h0m1s", func() {
		log.Println("run every 1 seconds")
	})
	go func() {
		cron.AddFunc("@every 0h0m2s", func() {
			log.Println("run every 2 seconds")
		})

	}()

	cancel := make(chan os.Signal)
	signal.Notify(cancel, os.Interrupt)
	select {
	case <-cancel:
		log.Println("dequeue job and run")
		cron.Stop()
	}
}
