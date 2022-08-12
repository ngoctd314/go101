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

type fileWriterConfig struct {
	filename     string
	age          int
	rorationTime string
}

// fileWriter with rotate feature comming from lumberjack
func newFileWriter(params fileWriterConfig) io.Writer {
	writer := &lumberjack.Logger{
		Filename:   params.filename,
		MaxSize:    0,
		MaxAge:     params.age,
		MaxBackups: 0,
		LocalTime:  false,
		Compress:   false,
	}

	return writer
}
