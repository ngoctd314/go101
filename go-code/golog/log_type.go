package golog

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogType log variation
type LogType string

// LogType enum
const (
	Sys LogType = "syslog" // log comming from system
)

func newSys(writer io.Writer, encoder encoder) Logger {
	core := zapcore.NewCore(
		encoder.zapEncoder(),
		zapcore.AddSync(writer),
		zapcore.InfoLevel, // level to enable
	)

	// >= Error => enable stacktrace
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel)).Sugar()
}
