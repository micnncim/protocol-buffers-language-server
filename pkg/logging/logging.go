package logging

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string, opts ...zap.Option) (*zap.Logger, error) {
	l, err := parseLogLevel(level)
	if err != nil {
		return nil, err
	}
	return newConfig(zap.NewAtomicLevelAt(l)).Build(opts...)
}

func newConfig(l zap.AtomicLevel) zap.Config {
	return zap.Config{
		Level:             l,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     logTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func logTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05"))
}

func parseLogLevel(levelStr string) (zapcore.Level, error) {
	switch strings.ToUpper(levelStr) {
	case zapcore.DebugLevel.CapitalString():
		return zapcore.DebugLevel, nil
	case zapcore.InfoLevel.CapitalString():
		return zapcore.InfoLevel, nil
	case zapcore.WarnLevel.CapitalString():
		return zapcore.WarnLevel, nil
	case zapcore.ErrorLevel.CapitalString():
		return zapcore.ErrorLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("undefined log level: %s", levelStr)
	}
}
