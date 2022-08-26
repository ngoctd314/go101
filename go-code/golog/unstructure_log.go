package golog

import (
	"io"
	"log"
)

func newUnStructureLog(writer io.Writer) logger {
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
func (*stdLogger) error(string) {
	panic("unimplemented")
}

// Errorw implements Logger
func (*stdLogger) errorw(string, ...any) {
	panic("unimplemented")
}

// Info implements Logger
func (*stdLogger) info(string) {
	panic("unimplemented")
}

// Infow implements Logger
func (*stdLogger) infow(string, ...any) {
	panic("unimplemented")
}

// Warn implements Logger
func (*stdLogger) warn(string) {
	panic("unimplemented")
}

// Warnw implements Logger
func (*stdLogger) warnw(string, ...any) {
	panic("unimplemented")
}
