package golog

import (
	"os"
)

// stdoutWriter write log to stdout (console)
type stdoutWriter struct{}

func newStdoutWriter() stdoutWriter {
	return stdoutWriter{}
}

// Write implements io.Writer
func (stdoutWriter) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}
