package zap

type Logger struct{}
type SugaredLogger struct{}

func L() *Logger {
	return &Logger{}
}
func S() *SugaredLogger {
	return &SugaredLogger{}
}

// Делаем методы с теми именами, которые проверяет levelSet
func (*Logger) Debug(msg string, fields ...any) {}
func (*Logger) Info(msg string, fields ...any)  {}
func (*Logger) Warn(msg string, fields ...any)  {}
func (*Logger) Error(msg string, fields ...any) {}

func (*SugaredLogger) Debug(msg string, args ...any) {}
func (*SugaredLogger) Info(msg string, args ...any)  {}
func (*SugaredLogger) Warn(msg string, args ...any)  {}
func (*SugaredLogger) Error(msg string, args ...any) {}
