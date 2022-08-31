package golog

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// sugared logger instance
func newZapLoggerCompose(zapEncoder zapcore.Encoder) loggerCompose {
	return func(writer io.Writer) logger {
		core := zapcore.NewCore(
			zapEncoder,
			zapcore.AddSync(writer),
			zapcore.InfoLevel, // level to enable
		)

		// >= Warnlevel => enable stacktrace
		zap := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))

		return &zapLogger{
			SugaredLogger: zap.Sugar(),
		}
	}
}

type zapEncoderOption func(*zapcore.EncoderConfig)

func newZapEncoderConfig(opts ...zapEncoderOption) zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		// TimeKey:    "ts",
		// NameKey:    "logger",
		// CallerKey:  "caller",
		// FunctionKey:    "",
		// StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     "\n",
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(l.CapitalString())
		},
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(time.Duration, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(caller.TrimmedPath())
		},
		EncodeName: func(name string, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(name)
		},
		// NewReflectedEncoder: func(io.Writer) zapcore.ReflectedEncoder {
		// },
		ConsoleSeparator: "",
	}

	// customize encoder config
	for _, opt := range opts {
		opt(&encoderConfig)
	}

	return encoderConfig
}

func zapEncoderWithTimeKey(timeKey string) zapEncoderOption {
	return func(encoderConfig *zapcore.EncoderConfig) {
		encoderConfig.TimeKey = timeKey
	}
}

func zapEncoderWithStacktraceKey(stacktraceKey string) zapEncoderOption {
	return func(encoderConfig *zapcore.EncoderConfig) {
		encoderConfig.StacktraceKey = stacktraceKey
	}
}

// structure log with zap
//
// zapLogger implement logger interface
type zapLogger struct {
	*zap.SugaredLogger
}

func (zap *zapLogger) info(msg string) {
	zap.Info(msg)
}

func (zap *zapLogger) infow(msg string, kv ...any) {
	zap.Infow(msg, kv...)
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
