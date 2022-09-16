package httpclient

// Response httpclient response
type Response struct {
	// Body is http response.body
	Body       []byte
	Err        error
	StatusCode int
}
