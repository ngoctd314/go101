package golog

import (
	"io"
	"os"

	"github.com/robfig/cron/v3"
)

type fileWriter struct {
	writer        io.Writer
	rotatePattern string
	cron          *cron.Cron
}

func newFileWriter(filePath string, opts ...func(*fileWriter)) fileWriter {
	_ = os.MkdirAll(filePath, 0644)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fwriter := fileWriter{
		writer: f,
	}

	return fwriter
}

// Write ...
func (f fileWriter) Write(p []byte) (n int, err error) {
	return f.writer.Write(p)
}
