package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 100)
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
}
