package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "cpu.prof", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "mem.prof", "write memory profile to `file`")

func main() {
	flag.Parse()
	go func() {
		http.HandleFunc("/cpu_profile", func(w http.ResponseWriter, r *http.Request) {
			if len(*cpuprofile) != 0 {
				f, err := os.Create(*cpuprofile)
				if err != nil {
					log.Fatal("could not create CPU profile: ", err)
				}
				defer f.Close()

				if err := pprof.StartCPUProfile(f); err != nil {
					log.Fatal("could not start CPU profile: ", err)
				}
				defer pprof.StopCPUProfile()
			}
			w.Write([]byte("write to cpu.profile"))
		})

		http.HandleFunc("/mem_profile", func(w http.ResponseWriter, r *http.Request) {
			if len(*memprofile) != 0 {
				f, err := os.Create(*memprofile)
				if err != nil {
					log.Fatal("cound not create memory profile: ", err)
				}
				defer f.Close()
				runtime.GC()
				if err := pprof.WriteHeapProfile(f); err != nil {
					log.Fatal("cound not write memory profile: ", err)
				}
			}
			w.Write([]byte("write to mem.profile"))
		})

		http.ListenAndServe(":6060", nil)
	}()

	go func() {
		str := make([]byte, 0)
		for {
			time.Sleep(time.Millisecond)
			tmp := make([]byte, 1000)
			str = append(str, tmp...)
		}
	}()

	cancel := make(chan os.Signal)
	signal.Notify(cancel, os.Interrupt)
	<-cancel
}
