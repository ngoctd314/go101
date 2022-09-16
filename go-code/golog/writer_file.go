package golog

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

// fileWriter write log to file
//
// fileWriter implement io.Writer interface
type fileWriter struct {
	filename string
	opts     []fileWriterOption
}

// fileWriter with rotate feature comming from lumberjack
func newFileWriter(filename string, opts ...fileWriterOption) fileWriter {
	return fileWriter{
		filename: filename,
		opts:     opts,
	}
}

// Write implements io.Writer
func (w fileWriter) Write(p []byte) (n int, err error) {
	writer := &lumberjack.Logger{
		Filename: w.filename,
	}
	for _, opt := range w.opts {
		opt(writer)
	}

	return writer.Write(p)
}

type fileWriterOption func(*lumberjack.Logger)

// withMaxAge set maxAge for logger file
func withMaxAge(maxAge int) fileWriterOption {
	return func(l *lumberjack.Logger) {
		l.MaxAge = maxAge
	}
}

func withMaxSize(maxSize int) fileWriterOption {
	return func(l *lumberjack.Logger) {
		l.MaxSize = maxSize
	}
}
