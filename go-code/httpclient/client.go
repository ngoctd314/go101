package httpclient

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/valyala/fasthttp"
)

// Client uses fasthttp as http client
type Client struct {
	lib *fasthttp.Client
}

// NewClient creates new httpclient and uses fasthttp as dependency
func NewClient(opts ...ClientOption) *Client {
	client := &Client{
		lib: &fasthttp.Client{
			NoDefaultUserAgentHeader: true,
			Dial: func(addr string) (net.Conn, error) {
				conn, err := net.Dial("tcp", addr)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return conn, nil
			},
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Do http request with fasthttp lib
//
// params:
// ctx: context propagation
// args: requirement parameter for execute http request
//
// return:
// response, statusCode, and error (if has)
func (c *Client) Do(ctx context.Context, args Args) (result Response) {
	var (
		httpRequest  = fasthttp.AcquireRequest()
		httpResponse = fasthttp.AcquireResponse()
		err          error
	)

	defer func() {
		fasthttp.ReleaseRequest(httpRequest)
		fasthttp.ReleaseResponse(httpResponse)

		if err := recover(); err != nil {
			result = Response{
				Body:       []byte{},
				Err:        err.(error),
				StatusCode: 400,
			}
		}
	}()

	// validate argument
	err = args.validate()
	if err != nil {
		panic(err)
	}

	httpRequest.SetRequestURI(args.URL)
	httpRequest.Header.SetMethod(args.Method)

	if args.Body != nil {
		httpRequest.SetBody(args.Body)
	}
	if args.Header != nil {
		for k, v := range args.Header {
			httpRequest.Header.Add(k, v)
		}
	}
	if args.Query != nil {
		for k, v := range args.Query {
			httpRequest.URI().QueryArgs().Set(k, v)
		}
	}
	if args.Timeout != 0 {
		err = c.lib.DoTimeout(httpRequest, httpResponse, args.Timeout)
	} else {
		err = c.lib.Do(httpRequest, httpResponse)
	}

	if err != nil {
		result = Response{
			Body:       []byte{},
			Err:        err,
			StatusCode: 400,
		}
		return
	}

	result = Response{
		Body:       httpResponse.Body(),
		Err:        err,
		StatusCode: 200,
	}

	return
}
