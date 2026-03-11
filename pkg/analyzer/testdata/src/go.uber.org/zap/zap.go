package zap

type Logger struct{}

func (l *Logger) Info(msg string, fields ...any)  {}
func (l *Logger) Error(msg string, fields ...any) {}
func (l *Logger) Warn(msg string, fields ...any)  {}
func (l *Logger) Debug(msg string, fields ...any) {}

func NewProduction() (*Logger, error) {
	return &Logger{}, nil
}

type Field struct{}

func String(key, val string) Field  { return Field{} }
func Int(key string, val int) Field { return Field{} }
