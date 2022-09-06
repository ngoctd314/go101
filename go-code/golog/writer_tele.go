package golog

import (
	"fmt"
	"io"
)

// teleWriter send log to tele
//
// teleWriter implement io.Writer interface
type teleWriter struct{}

func newTeleWriter() io.Writer {
	return &teleWriter{}
}

// Write implements io.Writer
func (*teleWriter) Write(p []byte) (n int, err error) {
	fmt.Println("RUN tele writer", string(p))
	return len(p), nil
}
