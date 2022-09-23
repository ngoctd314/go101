package httpclient

import "time"

// ClientOption implements optional pattern to configure Client object
type ClientOption func(*Client)

// WithMaxConns set maxinum number of connections per each host which may be established
//
// DefaultMaxConnsPerHost is used if not set
func WithMaxConns(conns int) ClientOption {
	return func(c *Client) {
		c.lib.MaxConnsPerHost = conns
	}
}

// WithDuration keep-alive connections are closed after this duration.
//
// By default connection duration is unlimited.
func WithDuration(duration time.Duration) ClientOption {
	return func(c *Client) {
		c.lib.MaxConnDuration = duration
	}
}

// WithTimeout ...
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.lib.ReadTimeout = timeout
	}
}
