package golog

type logger interface {
	info(string)
	infow(string, ...any)

	warn(string)
	warnw(string, ...any)

	error(string)
	errorw(string, ...any)
}

// LogType log variation
type LogType string

// LogType enum
const (
	Server LogType = "server-log" // log comming from system
)

// mapping LogType -> LogType instance
var _logMgmt = make(map[LogType]logger)

func init() {
	_logMgmt[Server] = newStructureLog(newStdoutWriter(), &zapEncoder{
		verbose: false,
	})
}

// Info log with info level
func Info(logType LogType, msg string) {
	log := _logMgmt[logType]

	log.infow(msg, "type", logType)
}

// Infow log with info level, with key-value pair as option
func Infow(logType LogType, msg string, kv ...any) {
	log := _logMgmt[logType]

	log.infow(msg, kv...)
}

// Warn log with warn level
func Warn(logType LogType, msg string) {
	log := _logMgmt[logType]

	log.warnw(msg, "type", logType)
}

// Warnw log with warn level, with key-value pair as option
func Warnw(logType LogType, msg string, kv ...any) {
	log := _logMgmt[logType]

	log.warnw(msg, kv...)
}

// Error log with error level
func Error(logType LogType, msg string) {
	log := _logMgmt[logType]

	log.errorw(msg, "type", logType)
}

// Errorw log with error level, with key-value pair as option
func Errorw(logType LogType, msg string, kv ...any) {
	log := _logMgmt[logType]

	log.errorw(msg, kv...)
}
