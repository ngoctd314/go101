package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ngoctd314/go101/go-code/httpclient"
)

type httpFetcher interface {
	Do(ctx context.Context, args httpclient.Args) httpclient.Response
}

type server struct {
	// httpclient .
	httpclient httpFetcher
}

func main() {
	srv := &server{
		httpclient: httpclient.NewClient(
			httpclient.WithDuration(time.Millisecond),
			httpclient.WithTimeout(time.Millisecond*100),
			httpclient.WithMaxConns(1000000),
		),
	}

	for i := 0; i < 2; i++ {
		go func(i int) {
			var u string
			if i%2 == 0 {
				u = "http://localhost:8080/api/fast"
			} else {
				u = "http://localhost:8081/api/fast"
			}

			resp := srv.httpclient.Do(context.TODO(), httpclient.Args{
				URL:     u,
				Method:  http.MethodGet,
				Body:    []byte{},
				Header:  map[string]string{},
				Query:   map[string]string{},
				Timeout: 0,
			})
			log.Println(string(resp.Body), resp.Err)
		}(i)
	}

	cancel := make(chan os.Signal)
	signal.Notify(cancel, os.Interrupt)
	select {
	case <-cancel:
		log.Println("Server is shutdown at: ", time.Now().String())
	}

}
