package httpclient

import (
	"errors"
	"strings"
	"time"
)

// Args parameter
type Args struct {
	URL     string
	Method  string
	Body    []byte
	Header  map[string]string
	Query   map[string]string
	Timeout time.Duration
}

func (args Args) validate() error {
	if len(strings.TrimSpace(args.URL)) == 0 {
		return errors.New("got: len(args.URL) = 0, want: len(args.URL) != 0")
	}
	if len(strings.TrimSpace(args.Method)) == 0 {
		return errors.New("got: len(args.Method) = 0, want: len(args.Method) != 0")
	}

	return nil
}
