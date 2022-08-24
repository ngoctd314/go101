package golog

// Logger ...
type Logger interface {
	Info(string)
	Infow(string, ...any)

	Warn(string)
	Warnw(string, ...any)

	Error(string)
	Errorw(string, ...any)
}

var (
	logMgmt = make(map[LogType]Logger)
)

func init() {
	logMgmt[Sys] = newStructureLog(newStdoutWriter(), &zapEncoder{
		verbose: false,
	})
}

// Info log with info level
func Info(logType LogType, msg string) {
	log := logMgmt[logType]

	log.Infow(msg, "type", logType)
}

// Infow log with info level, with key-value pair as option
func Infow(logType LogType, msg string, kv ...any) {
	log := logMgmt[logType]

	log.Infow(msg, kv...)
}

// Warn log with warn level
func Warn(logType LogType, msg string) {
	log := logMgmt[logType]

	log.Warnw(msg, "type", logType)
}

// Warnw log with warn level, with key-value pair as option
func Warnw(logType LogType, msg string, kv ...any) {
	log := logMgmt[logType]

	log.Warnw(msg, kv...)
}

// Error log with error level
func Error(logType LogType, msg string) {
	log := logMgmt[logType]

	log.Errorw(msg, "type", logType)
}

// Errorw log with error level, with key-value pair as option
func Errorw(logType LogType, msg string, kv ...any) {
	log := logMgmt[logType]

	log.Errorw(msg, kv...)
}
