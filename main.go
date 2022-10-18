package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ngoctd314/go101/go-code/httpclient"
	"github.com/ngoctd314/go101/go-code/workerpool"
)

var hclient = httpclient.NewClient(httpclient.WithMaxConns(10_000))

type httpJob struct {
	res chan httpclient.Response
	wg  *sync.WaitGroup
}

func (h httpJob) Do() error {
	defer h.wg.Done()

	res := hclient.Do(context.TODO(), httpclient.Args{
		URL:    "http://localhost:8080",
		Method: http.MethodGet,
	})

	h.res <- res

	return nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		wg := &sync.WaitGroup{}
		var n = 10
		res := make(chan httpclient.Response, n)
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				job := httpJob{
					res: res,
					wg:  wg,
				}
				workerpool.EnQueue(job)
			}()
		}
		wg.Wait()

		since := fmt.Sprint(time.Since(now))
		w.Write([]byte(since))
	})

	http.ListenAndServe(":8081", nil)
}
