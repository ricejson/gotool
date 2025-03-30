package logx

import "go.uber.org/zap"

// ZapLogger 封装zap的实现
type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() Logger {
	logger, _ := zap.NewDevelopment()
	return &ZapLogger{
		logger: logger,
	}
}

func (l *ZapLogger) toArgs(fields []Field) []zap.Field {
	var args = make([]zap.Field, 0, len(fields))
	for _, field := range fields {
		args = append(args, zap.Any(field.Key, field.Value))
	}
	return args
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, l.toArgs(fields)...)
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, l.toArgs(fields)...)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, l.toArgs(fields)...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, l.toArgs(fields)...)
}
