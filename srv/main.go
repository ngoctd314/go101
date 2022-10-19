package main

import (
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// n := rand.Intn(10000)
		// if n%2 == 0 && n%3 == 0 {
		// 	time.Sleep(time.Millisecond * 50)
		// }
		time.Sleep(time.Millisecond * 50)
		data := make([]byte, 10000)
		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
