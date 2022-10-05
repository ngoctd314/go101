package golog

import "os"

type fileWriter struct {
	f *os.File
}

func newFileWriter()

type fileCreator struct{}
