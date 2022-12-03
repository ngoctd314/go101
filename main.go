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
		httpclient.WithMaxKeepAliveIdleDuration(time.Second*2),
	)

	ctx := context.Background()

	resp := hc.DoMany(ctx,
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

	resp1 := hc.DoMany(ctx,
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
	)

	for v := range resp1 {
		if v.Err != nil {
			log.Panic(v.Err)
		}
		fmt.Println(string(v.Body))
	}

	time.Sleep(time.Second * 1)

	resp2 := hc.DoMany(ctx,
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
	)

	for v := range resp2 {
		if v.Err != nil {
			log.Panic(v.Err)
		}
		fmt.Println(string(v.Body))
	}
	time.Sleep(time.Second * 1)

	resp3 := hc.DoMany(ctx,
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
	)

	for v := range resp3 {
		if v.Err != nil {
			log.Panic(v.Err)
		}
		fmt.Println(string(v.Body))
	}
	time.Sleep(time.Second * 1)

	resp4 := hc.DoMany(ctx,
		httpclient.Args{
			URL:    "http://localhost:8080",
			Method: http.MethodGet,
		},
	)

	for v := range resp4 {
		if v.Err != nil {
			log.Panic(v.Err)
		}
		fmt.Println(string(v.Body))
	}
}
