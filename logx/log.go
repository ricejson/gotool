package logx

// Logger 统一日志接口
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
}

func String(key string, val string) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}

func Error(err error) Field {
	return Field{
		Key:   "error",
		Value: err,
	}
}

func Int(key string, val int) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}

type Field struct {
	Key   string
	Value any
}
