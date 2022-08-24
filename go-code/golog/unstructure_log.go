package golog

import (
	"io"
	"log"
)

func newUnStructureLog(writer io.Writer) Logger {
	stdLog := &stdLogger{
		Logger: new(log.Logger),
	}
	stdLog.SetOutput(writer)
	stdLog.SetFlags(0)

	return stdLog
}

type stdLogger struct {
	*log.Logger
}

// Error implements Logger
func (*stdLogger) Error(string) {
	panic("unimplemented")
}

// Errorw implements Logger
func (*stdLogger) Errorw(string, ...any) {
	panic("unimplemented")
}

// Info implements Logger
func (*stdLogger) Info(string) {
	panic("unimplemented")
}

// Infow implements Logger
func (*stdLogger) Infow(string, ...any) {
	panic("unimplemented")
}

// Warn implements Logger
func (*stdLogger) Warn(string) {
	panic("unimplemented")
}

// Warnw implements Logger
func (*stdLogger) Warnw(string, ...any) {
	panic("unimplemented")
}
