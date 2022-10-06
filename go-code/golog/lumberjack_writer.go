package golog

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

// lumberjackWriter write log to file
//
// fileWriter implement io.Writer interface
type lumberjackWriter struct {
	writer io.Writer
}

// newLumberJackWriter with rotate feature comming from lumberjack
func newLumberJackWriter(filePath string, opts ...func(*lumberjack.Logger)) lumberjackWriter {
	writer := &lumberjack.Logger{
		Filename: filePath,
	}
	for _, opt := range opts {
		opt(writer)
	}

	return lumberjackWriter{
		writer: writer,
	}
}

// Write implements io.Writer
func (w lumberjackWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

// withLumberJackMaxAge set maxAge for logger file
func withLumberJackMaxAge(maxAge int) func(*lumberjack.Logger) {
	return func(l *lumberjack.Logger) {
		l.MaxAge = maxAge
	}
}

// withLumberJackMaxSize set maxSize for logger
func withLumberJackMaxSize(maxSize int) func(*lumberjack.Logger) {
	return func(l *lumberjack.Logger) {
		l.MaxSize = maxSize
	}
}
