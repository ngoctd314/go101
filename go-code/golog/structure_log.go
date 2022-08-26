package golog

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newStructureLog(writer io.Writer, enc encoder) logger {
	core := zapcore.NewCore(
		enc.encode().(zapcore.Encoder),
		zapcore.AddSync(writer),
		zapcore.InfoLevel, // level to enable
	)

	// >= Error => enable stacktrace
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel)).Sugar()
	return newZapLogger(zapLogger)
}

// structure log with zap (sugger)
type zapLogger struct {
	*zap.SugaredLogger
}

func newZapLogger(sugarLogger *zap.SugaredLogger) logger {
	return &zapLogger{
		SugaredLogger: sugarLogger,
	}
}

func (zap *zapLogger) info(msg string) {
	zap.Info(msg)
}

func (zap *zapLogger) infow(msg string, kv ...any) {
	zap.Warnw(msg, kv...)
}

func (zap *zapLogger) warn(msg string) {
	zap.Warn(msg)
}

func (zap *zapLogger) warnw(msg string, kv ...any) {
	zap.Warnw(msg, kv...)
}
func (zap *zapLogger) error(msg string) {
	zap.Error(msg)
}

func (zap *zapLogger) errorw(msg string, kv ...any) {
	zap.Errorw(msg, kv...)
}
