package logging

type StdOutLogger struct {
}

func (l *StdOutLogger) Log(event LogEvent) {
	switch event.Level {
		case InfoLevel:
			println(LogLevelToString(event.Level) + ": " + event.Message)
		case DebugLevel:
			println("")
		case ErrorLevel:
			println(LogLevelToString(event.Level) + ": " + event.Message + event.Error.Error())
		case FatalLevel:
			println("")
	}
}

func NewStdOutLogger() *StdOutLogger {
	return &StdOutLogger{}
}
