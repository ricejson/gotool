package logx

import "log/slog"

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() Logger {
	return &SlogLogger{
		logger: slog.Default(),
	}
}

func (l *SlogLogger) toArgs(fields []Field) []any {
	var args = make([]any, 0, 2*len(fields))
	for _, field := range fields {
		args = append(args, field.Key, field.Value)
	}
	return args
}

func (l *SlogLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, l.toArgs(fields)...)
}

func (l *SlogLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, l.toArgs(fields)...)
}

func (l *SlogLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, l.toArgs(fields)...)
}

func (l *SlogLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, l.toArgs(fields)...)
}
