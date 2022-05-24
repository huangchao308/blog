package logger

import (
	"context"

	"github.com/huangchao308/blog/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fields map[string]interface{}

type Logger struct {
	ZapLogger *zap.Logger
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
			MessageKey:  settings.EncoderConfig.MessageKey,
			LevelKey:    settings.EncoderConfig.LevelKey,
			EncodeLevel: zapcore.LowercaseColorLevelEncoder,
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

func (l *Logger) format(ctx context.Context) string {
	return ""
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	defer l.ZapLogger.Sync()
	l.ZapLogger.Sugar().Debugf(msg, args...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.ZapLogger.Sugar().Infof(msg, args...)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.ZapLogger.Sugar().Warnf(msg, args...)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.ZapLogger.Sugar().Errorf(msg, args...)
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.ZapLogger.Sugar().Fatalf(msg, args...)
}
