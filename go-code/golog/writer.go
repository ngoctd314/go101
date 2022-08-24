package golog

import (
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// writer: how to log write
// writer implement io.Writer interface
//
// stdout
// file

func newStdoutWriter() io.Writer {
	return os.Stdout
}

// fileWriter with rotate feature comming from lumberjack
func newFileWriter(filename string, opts ...fileWriterOption) io.Writer {
	writer := &lumberjack.Logger{
		Filename: filename,
	}

	return writer
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
