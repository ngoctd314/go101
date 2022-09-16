package httpclient

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
