package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.Header().Add("Content-Disposition", `attachment; filename="data.txt"`)
		w.Write([]byte("HEllo WORLD"))
	})

	http.ListenAndServe(":8081", nil)

}
