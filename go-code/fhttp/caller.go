package fhttp

import (
	"context"
	"time"

	"github.com/valyala/fasthttp"
)

// Response fhttp response
type Response struct {
	Body       []byte
	Err        error
	StatusCode int
}

// Args parameter
type Args struct {
	URL     string
	Method  string
	Body    []byte
	Header  map[string]string
	Query   map[string]string
	Timeout time.Duration
}

// Caller exec fetch http request
type Caller interface {
	Do(ctx context.Context, args Args) Response
}

// NewClient return new fasthttp instance
func NewClient() Caller {
	return &client{
		Client: &fasthttp.Client{
			// Name:                     "",
			NoDefaultUserAgentHeader: true,
			// Dial: func(addr string) (net.Conn, error) {
			// },
			// DialDualStack:                 false,
			// TLSConfig:                     &tls.Config{},
			MaxConnsPerHost: 5000,
			// MaxIdleConnDuration:           0,
			// MaxConnDuration:               0,
			// MaxIdemponentCallAttempts:     0,
			// ReadBufferSize:                0,
			// WriteBufferSize:               0,
			// ReadTimeout:                   0,
			// WriteTimeout:                  0,
			// MaxResponseBodySize:           0,
			// DisableHeaderNamesNormalizing: false,
			// DisablePathNormalizing:        false,
			// MaxConnWaitTimeout:            0,
			// RetryIf: func(request *fasthttp.Request) bool {
			// },
			// ConnPoolStrategy: 0,
			// ConfigureClient: func(hc *fasthttp.HostClient) error {
			// },
		},
	}
}

type client struct {
	*fasthttp.Client
}

func (c *client) Do(ctx context.Context, args Args) Response {
	var (
		req  = fasthttp.AcquireRequest()
		resp = fasthttp.AcquireResponse()
		err  error
	)

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	// inject args
	if args.URL == "" {
		panic("missing args.URL")
	}
	req.SetRequestURI(args.URL)

	if args.Method == "" {
		panic("missing args.Method")
	}
	req.Header.SetMethod(args.Method)

	if args.Body != nil {
		req.SetBody(args.Body)
	}
	if args.Header != nil {
		for k, v := range args.Header {
			req.Header.Add(k, v)
		}
	}
	if args.Query != nil {
		for k, v := range args.Query {
			req.URI().QueryArgs().Set(k, v)
		}
	}
	if args.Timeout != 0 {
		err = c.Client.DoTimeout(req, resp, args.Timeout)
	} else {
		err = c.Client.Do(req, resp)
	}

	if err != nil {
		return Response{
			Body:       []byte{},
			Err:        err,
			StatusCode: 400,
		}
	}

	return Response{
		Body:       resp.Body(),
		Err:        err,
		StatusCode: 200,
	}
}
