package logging

type Logger interface {
	Log(event LogEvent)
}

type LogEvent struct {
	Level   LogLevel
	Message string
}

type LogLevel string

const (
	InfoLevel  LogLevel = "INFO"
	DebugLevel LogLevel = "DEBUG"
	ErrorLevel LogLevel = "ERROR"
	FatalLevel LogLevel = "FATAL"
)

func LogLevelToString(level LogLevel) string {
	return string(level)
}

func NewLogEvent(level LogLevel, message string) LogEvent {
	return LogEvent{Level: level, Message: message}
}
