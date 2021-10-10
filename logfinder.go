package logging

// LogFinder defines methods for finding named loggers.
// All methods should be safe for concurrent use.
type LogFinder interface {
	// Find returns the Logger instance with the specific name.
	// If the name does not exist, returns nil.
	Find(name string) Logger
}

// NewSingleLoggerLogFinder creates a new instance of SingleLoggerLogFinder.
func NewSingleLoggerLogFinder(logger Logger) LogFinder {
	return &SingleLoggerLogFinder{logger}
}

// SingleLoggerLogFinder is a LogFinder whose Find() will always return the same Logger.
type SingleLoggerLogFinder struct {
	Logger Logger
}

var _ LogFinder = (*SingleLoggerLogFinder)(nil)

func (f SingleLoggerLogFinder) Find(name string) Logger {
	return f.Logger
}
