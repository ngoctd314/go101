package golog

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type encoder interface {
	zapEncoder() zapcore.Encoder
}

// encoder: how to log represent (format)
//
// json
// text

type jsonEncoder struct {
	verbose bool
}

func (e *jsonEncoder) zapEncoder() zapcore.Encoder {
	defaultKey := struct {
		msg           string
		levelKey      string
		stacktraceKey string
	}{
		msg: "msg",
	}

	if e.verbose {
		defaultKey.levelKey = "level"
		defaultKey.stacktraceKey = "stacktrace"
	}

	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     defaultKey.msg,
		LevelKey:       defaultKey.levelKey,
		TimeKey:        "", // use es ts
		NameKey:        "logger",
		CallerKey:      "",
		FunctionKey:    "",
		StacktraceKey:  defaultKey.stacktraceKey,
		SkipLineEnding: false,
		LineEnding:     "\n",
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(l.String())
		},
		// EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// 	enc.AppendString(t.Format("2006-01-02 15:04:05"))
		// },
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendFloat64(float64(d) / float64(time.Second))
		},
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(caller.TrimmedPath())
		},
		EncodeName: func(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(loggerName)
		},
		// NewReflectedEncoder: func(io.Writer) zapcore.ReflectedEncoder {
		// },
		ConsoleSeparator: "",
	})
}
