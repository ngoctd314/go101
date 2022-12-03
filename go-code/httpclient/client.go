package httpclient

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/valyala/fasthttp"
)

// Client uses fasthttp as http client
type Client struct {
	lib *fasthttp.Client
}

var count = 0
var mu sync.Mutex

// NewClient creates new httpclient and uses fasthttp as dependency
func NewClient(opts ...ClientOption) *Client {
	client := &Client{
		lib: &fasthttp.Client{
			NoDefaultUserAgentHeader: true,
			Dial: func(addr string) (net.Conn, error) {
				mu.Lock()
				count++
				log.Println("call dial times: ", count)
				mu.Unlock()
				return net.Dial("tcp", addr)
			},
			MaxIdemponentCallAttempts: 1,
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
func (c *Client) Do(ctx context.Context, args Args) Response {
	var (
		httpRequest  = fasthttp.AcquireRequest()
		httpResponse = fasthttp.AcquireResponse()
		err          error
	)

	defer func() {
		// release fasthttp resource
		fasthttp.ReleaseResponse(httpResponse)
		fasthttp.ReleaseRequest(httpRequest)
	}()

	// validate argument
	err = args.validate()
	if err != nil {
		// invalid argument case, return error with status code 400
		return errResponse(err)
	}

	httpRequest.SetRequestURI(args.URL)
	httpRequest.Header.SetMethod(args.Method)

	// set body in case POST method
	if args.Method == http.MethodPost && args.Body != nil {
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

	// exec fasthttp request
	if args.Timeout != 0 {
		err = c.lib.DoTimeout(httpRequest, httpResponse, args.Timeout)
	} else {
		err = c.lib.Do(httpRequest, httpResponse)
	}

	if err != nil {
		// exec http error, return error with status code 400
		return errResponse(err)
	}

	return Response{
		Body:       httpResponse.Body(),
		Err:        err,
		StatusCode: 200,
	}
}

func errResponse(err error) Response {
	return Response{
		Body:       []byte{},
		Err:        err,
		StatusCode: 400,
	}
}

// DoMany request concurrency
func (c *Client) DoMany(ctx context.Context, args ...Args) <-chan Response {
	var (
		outbound = make(chan Response, len(args))
		wg       sync.WaitGroup
	)

	wg.Add(len(args))
	for _, arg := range args {
		go func(arg Args) {
			defer wg.Done()
			outbound <- c.Do(ctx, arg)
		}(arg)
	}

	go func() {
		// pipeline principle
		// stages close their outbound channels when all the send operations are done
		wg.Wait()
		close(outbound)
	}()

	return outbound
}
