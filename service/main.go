package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET: ", r.URL.RequestURI())
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}
