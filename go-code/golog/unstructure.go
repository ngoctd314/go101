package golog

import (
	"fmt"
	"io"
	"log"
)

func newUnStructureLog(writer io.Writer) *stdLogger {
	stdLog := &stdLogger{
		Logger: new(log.Logger),
	}
	stdLog.SetOutput(writer)
	stdLog.SetFlags(log.Ldate | log.Ltime)

	return stdLog
}

type stdLogger struct {
	*log.Logger
}

// Info implements Logger
func (std *stdLogger) info(msg string) {
	std.Logger.Println(msg)
}

// Infow implements Logger
func (std *stdLogger) infow(msg string, kv ...any) {
	if len(kv)%2 != 0 {
		kv = append(kv, "")
	}

	for i := 0; i < len(kv); i = i + 2 {
		msg += fmt.Sprintf("\t%s:%s", kv[i], kv[i+1])
	}

	std.Logger.Println(msg)
}

// Error implements Logger
func (std *stdLogger) error(msg string) {
	panic("unimplemented")
}

// Errorw implements Logger
func (std *stdLogger) errorw(msg string, kv ...any) {
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
