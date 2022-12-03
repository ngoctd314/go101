package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ngoctd314/go101/go-code/httpclient"
)

func main() {
	hc := httpclient.NewClient(
		httpclient.WithMaxConns(100),
		httpclient.WithMaxKeepAliveIdleDuration(time.Second*10),
		httpclient.WithDialFunc(1),
	)

	ctx := context.Background()

	resp := hc.DoMany(ctx,
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
	)

	for v := range resp {
		if v.Err != nil {
			log.Panic(v.Err)
		}
		fmt.Println(string(v.Body))
	}
	time.Sleep(time.Second * 1)

}
