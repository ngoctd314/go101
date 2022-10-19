package main

import (
	"fmt"
	"sync"
	"time"

	_ "net/http/pprof"

	"github.com/ngoctd314/go101/go-code/httpclient"
	"github.com/ngoctd314/go101/go-code/workerpool"
)

func main() {
	fn1()
}

var hc = httpclient.NewClient(httpclient.WithMaxConns(10_000))

type httpjob struct {
	wg  *sync.WaitGroup
	res chan int
}

func (h httpjob) Do() error {
	defer h.wg.Done()
	time.Sleep(time.Second)
	h.res <- 1
	return nil
}

var n = 1000

func fn() {

	wg := &sync.WaitGroup{}
	res := make(chan int, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		job := httpjob{
			wg:  wg,
			res: res,
		}
		go job.Do()
	}

	wg.Wait()
	close(res)
	s := 0
	for r := range res {
		s += r
	}
	fmt.Println(s)
}

func fn1() {
	wg := &sync.WaitGroup{}
	res := make(chan int, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		job := httpjob{
			wg:  wg,
			res: res,
		}
		workerpool.EnQueue(job)
	}

	wg.Wait()
	close(res)
	s := 0
	for r := range res {
		s += r
	}
	fmt.Println(s)
}
