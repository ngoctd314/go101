package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	http.HandleFunc("/api/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Minute * 10)
		w.Write([]byte("hello world"))
	})

	http.HandleFunc("/api/fast", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world srv 0"))
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Println("ListenAndServe error: ", err)
		}
	}()

	log.Println("listen and server on port: 8080")
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	select {
	case <-interrupt:
		log.Println("force quit!")
	}
}
