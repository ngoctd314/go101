package golog

import (
	"io"

	"go.uber.org/zap/zapcore"
)

type logger interface {
	info(string)
	infow(string, ...any)

	warn(string)
	warnw(string, ...any)

	error(string)
	errorw(string, ...any)
}

type loggerCompose func(writer io.Writer) logger

// LogType log variation
type LogType string

// LogType enum
const (
	Server LogType = "server-log" // log comming from system
	Test   LogType = "test-log"
)

// mapping LogType -> LogType instance
var _logMgmt = make(map[LogType]logger)

func init() {
	// write
	stdoutWriter := newStdoutWriter()

	// structure log with zap
	zapEncoderConfig := newZapEncoderConfig(zapEncoderWithTimeKey("ts"), zapEncoderWithStacktraceKey("stacktrace"))
	zapConsoleEncoder := zapcore.NewJSONEncoder(zapEncoderConfig)
	zapLoggerCompose := newZapLoggerCompose(zapConsoleEncoder)

	_logMgmt[Server] = zapLoggerCompose(stdoutWriter)
	_logMgmt[Test] = zapLoggerCompose(newTeleWriter())
}

// Info log with info level
func Info(logType LogType, msg string) {
	_logMgmt[logType].infow(msg, "type", logType)
}

// Infow log with info level, with key-value pair as option
func Infow(logType LogType, msg string, kv ...any) {
	kv = append(kv, "type", logType)
	_logMgmt[logType].infow(msg, kv...)
}

// Warn log with warn level
func Warn(logType LogType, msg string) {
	_logMgmt[logType].warnw(msg, "type", logType)
}

// Warnw log with warn level, with key-value pair as option
func Warnw(logType LogType, msg string, kv ...any) {
	kv = append(kv, "type", logType)
	_logMgmt[logType].warnw(msg, kv...)
}

// Error log with error level
func Error(logType LogType, msg string) {
	_logMgmt[logType].errorw(msg, "type", logType)
}

// Errorw log with error level, with key-value pair as option
func Errorw(logType LogType, msg string, kv ...any) {
	kv = append(kv, "type", logType)
	_logMgmt[logType].errorw(msg, kv...)
}
