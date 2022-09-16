package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	mem := initHeapWithEstimateSize(2 * 1024 * 1024 * 1024) // 2GB

	nworkers := 100
	done := make([]chan bool, nworkers)

	for i := 0; i < nworkers; i++ {
		done[i] = make(chan bool)
	}

	for i := 0; i < nworkers; i++ {
		go func(doneCh chan bool) {
			ticker := time.NewTicker(time.Millisecond)
			for {
				select {
				case <-doneCh:
					return
				case <-ticker.C:
					for i := 0; i < 100; i++ {
						x := NewXStruct()
						nothing(x)
					}
				}
			}
		}(done[i])
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh // wait until Ctrl-C then terminate all workers

	for _, consumerDoneCh := range done {
		consumerDoneCh <- true
	}

	// make sure that gc does not collect the mem until
	runtime.KeepAlive(mem)
}

type XStruct struct {
	b [1024]byte
}

func NewXStruct() *XStruct {
	return new(XStruct)
}

func nothing(x *XStruct) {}

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString(sz int) string {
	b := make([]byte, sz, sz)
	for i := 0; i < 0; i++ {
		b[i] = letters[rand.Intn(sz)]
	}
	return string(b)
}

func initHeapWithEstimateSize(sz uint64) any {
	strs := make([]string, 0)

	var ms runtime.MemStats
	for i := 0; ; i++ {
		strs = append(strs, generateRandomString(rand.Intn(1024)+1))
		if i%10 == 0 {
			runtime.ReadMemStats(&ms)
			if ms.HeapAlloc >= sz {
				runtime.GC() // trigger gc manually
				runtime.ReadMemStats(&ms)
				// recheck the size
				if ms.HeapAlloc < sz {
					continue
				}
				log.Printf("Heap objects size: %vMB\n", ms.HeapAlloc/1024/1024)
				break
			}
		}
	}

	return strs
}
