package golog

import (
	"io"
	"log"
)

func newUnStructureLog(writer io.Writer) Logger {
	return newStdLoggerAdapter()
}

type stdLoggerAdapter struct {
}

func newStdLoggerAdapter() Logger {
	return &stdLoggerAdapter{}
}

// Info implements Logger
func (*stdLoggerAdapter) Info(msg string) {
	log.Println(msg)
}

// Infow implements Logger
func (*stdLoggerAdapter) Infow(msg string, kv ...any) {
	log.Println("NOOP")
}

// Error implements Logger
func (*stdLoggerAdapter) Error(string) {
	log.Println("NOOP")
}

// Errorw implements Logger
func (*stdLoggerAdapter) Errorw(string, ...any) {
	log.Println("NOOP")
}

// Warn implements Logger
func (*stdLoggerAdapter) Warn(string) {
	log.Println("NOOP")
}

// Warnw implements Logger
func (*stdLoggerAdapter) Warnw(string, ...any) {
	log.Println("NOOP")
}
