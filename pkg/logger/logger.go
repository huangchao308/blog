package logger

import (
	"context"
	"time"

	"github.com/huangchao308/blog/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fields map[string]interface{}
type Logger struct {
	ZapLogger *zap.Logger
	ctx       context.Context
}

func NewLogger(settings *setting.LogSettingS) (*Logger, error) {
	zapLevel, err := zap.ParseAtomicLevel(settings.Level)
	if err != nil {
		return nil, err
	}
	cfg := &zap.Config{
		Level:            zapLevel,
		Encoding:         settings.Encoding,
		OutputPaths:      settings.OutputPaths,
		ErrorOutputPaths: settings.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       settings.EncoderConfig.MessageKey,
			LevelKey:         settings.EncoderConfig.LevelKey,
			EncodeLevel:      zapcore.LowercaseColorLevelEncoder,
			ConsoleSeparator: "|",
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return &Logger{
		ZapLogger: logger,
	}, nil
}

func (l *Logger) format(ctx context.Context, msg string) string {
	now := time.Now().Format("2006-01-02 15:04:05")
	return now + "|" + msg
}

func (l *Logger) log(level string, msg string, args ...interface{}) {
	defer l.ZapLogger.Sync()
	msg = l.format(l.ctx, msg)
	switch level {
	case "debug":
		l.ZapLogger.Sugar().Debugf(msg, args...)
	case "info":
		l.ZapLogger.Sugar().Infof(msg, args...)
	case "warn":
		l.ZapLogger.Sugar().Warnf(msg, args...)
	case "error":
		l.ZapLogger.Sugar().Errorf(msg, args...)
	case "fatal":
		l.ZapLogger.Sugar().Fatalf(msg, args...)
	}
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.log("debug", msg, args...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.log("info", msg, args...)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.log("warn", msg, args...)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.log("error", msg, args...)
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.log("fatal", msg, args...)
}
