package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 5)
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8001", nil)
}
