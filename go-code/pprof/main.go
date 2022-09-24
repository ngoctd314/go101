package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()
	var wg sync.WaitGroup
	wg.Add(1)

	wg.Wait()
}
