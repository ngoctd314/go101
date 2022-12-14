package httpclient

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

// ClientOption implements optional pattern to configure Client object
type ClientOption func(*Client)

// WithMaxConns set maximum number of connections per each host which may be established
//
// DefaultMaxConnsPerHost is used if not set
func WithMaxConns(conns int) ClientOption {
	return func(c *Client) {
		c.lib.MaxConnsPerHost = conns
	}
}

// WithMaxKeepAliveIdleDuration Idle keep-alive connections are closed after this duration
//
// By default idle connections are closed after 10s
func WithMaxKeepAliveIdleDuration(duration time.Duration) ClientOption {
	return func(c *Client) {
		c.lib.MaxIdleConnDuration = duration
	}
}

// WithKeepAliveIntervalDuration keep-alive connections are closed after this duration.
//
// By default connection duration is unlimited.
func WithKeepAliveIntervalDuration(duration time.Duration) ClientOption {
	return func(c *Client) {
		c.lib.MaxConnDuration = duration
	}
}

// WithReadTimeout sets maximum duration for full response reading (including body).
//
// By default response read timeout is unlimited
func WithReadTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.lib.ReadTimeout = timeout
	}
}

// WithRetryTimes sets maximum number of retry times when connection is keep-alive or timeout
//
// By default fasthttp retrys 5 times
func WithRetryTimes(num int) ClientOption {
	return func(c *Client) {
		c.lib.MaxIdemponentCallAttempts = num
	}
}

// RetryIf ..
type RetryIf interface {
	Retry() bool
}

// WithRetryFunc controls whether a retry should be attempted after an error
//
func WithRetryFunc(r RetryIf) ClientOption {
	return func(c *Client) {
		c.lib.RetryIf = func(*fasthttp.Request) bool {
			return r.Retry()
		}
	}
}

// WithDialFunc custom concurrency amount
func WithDialFunc(concurrency int) ClientOption {
	return func(c *Client) {
		// fn := func(addr string) (net.Conn, error) {
		// 	mu.Lock()
		// 	count++
		// 	log.Println("call dial times: ", count)
		// 	mu.Unlock()
		// 	return net.Dial("tcp", addr)
		// }
		// c.lib.Dial = fn
		c.lib.Dial = func() fasthttp.DialFunc {
			log.Println("dial")
			return (&fasthttp.TCPDialer{
				Concurrency:      0,
				DNSCacheDuration: time.Hour, // cache DSN on 1 hour
			}).Dial
		}()
	}
}
