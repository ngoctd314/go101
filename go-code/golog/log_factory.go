package golog

import "sync"

// log_factory: combine writer, encoder and lib => golog
//

// Logger ...
type Logger interface {
	Info(...any)
	Infow(string, ...any)

	Warn(...any)
	Warnw(string, ...any)

	Error(...any)
	Errorw(string, ...any)
}

var (
	logMgmt = make(map[LogType]Logger)
	rwm     sync.RWMutex
)

func init() {
	logMgmt[Sys] = newSys(newStdoutWriter(), &jsonEncoder{
		verbose: false,
	})
}

// Info log with info level
func Info(logType LogType, msg string) {
	rwm.RLock()
	log := logMgmt[logType]
	rwm.RUnlock()

	log.Infow(msg, "type", logType)
}

// Infow log with info level, with key-value pair as option
func Infow(logType LogType, msg string, kv ...any) {
	rwm.RLock()
	log := logMgmt[logType]
	rwm.RUnlock()

	log.Infow(msg, kv...)
}

// Warn log with warn level
func Warn(logType LogType, msg string) {
	rwm.RLock()
	log := logMgmt[logType]
	rwm.RUnlock()

	log.Warnw(msg, "type", logType)
}

// Warnw log with warn level, with key-value pair as option
func Warnw(logType LogType, msg string, kv ...any) {
	rwm.RLock()
	log := logMgmt[logType]
	rwm.RUnlock()

	log.Warnw(msg, kv...)
}

// Error log with error level
func Error(logType LogType, msg string) {
	rwm.RLock()
	log := logMgmt[logType]
	rwm.RUnlock()

	log.Errorw(msg, "type", logType)
}

// Errorw log with error level, with key-value pair as option
func Errorw(logType LogType, msg string, kv ...any) {
	rwm.RLock()
	log := logMgmt[logType]
	rwm.RUnlock()

	log.Errorw(msg, kv...)
}
