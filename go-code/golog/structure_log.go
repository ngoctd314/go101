package golog

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newStructureLog(writer io.Writer, enc encoder) Logger {
	core := zapcore.NewCore(
		enc.encode().(zapcore.Encoder),
		zapcore.AddSync(writer),
		zapcore.InfoLevel, // level to enable
	)

	// >= Error => enable stacktrace
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel)).Sugar()
	return newZapAdapter(zapLogger)
}

// zapAdapter implement Logger interface
type zapLoggerAdapter struct {
	*zap.SugaredLogger
}

func newZapAdapter(zapLogger *zap.SugaredLogger) Logger {
	return &zapLoggerAdapter{
		SugaredLogger: zapLogger,
	}
}

func (adapter *zapLoggerAdapter) Info(msg string) {
	adapter.Info(msg)
}

func (adapter *zapLoggerAdapter) Infow(msg string, kv ...any) {
	adapter.Warnw(msg, kv...)
}

func (adapter *zapLoggerAdapter) Warn(msg string) {
	adapter.Warn(msg)
}

func (adapter *zapLoggerAdapter) Warnw(msg string, kv ...any) {
	adapter.Warnw(msg, kv...)
}
func (adapter *zapLoggerAdapter) Error(msg string) {
	adapter.Error(msg)
}

func (adapter *zapLoggerAdapter) Errorw(msg string, kv ...any) {
	adapter.Errorw(msg, kv...)
}
