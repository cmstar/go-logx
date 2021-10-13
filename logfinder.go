package logging

// LogFinder defines methods for finding named loggers.
// All methods should be safe for concurrent use.
type LogFinder interface {
	// Find returns the Logger instance with the specific name.
	// If the name cannot be found, returns nil.
	Find(name string) Logger
}

// NewSingleLoggerLogFinder creates a new instance of LogFinder, whose Find() will always
// return the same Logger.
// It is useful when there won't be more than one Logger, and a function receives a LogFinder.
// The logger can be nil.
func NewSingleLoggerLogFinder(logger Logger) LogFinder {
	return &singleLoggerLogFinder{logger}
}

type singleLoggerLogFinder struct {
	l Logger // The logger to return.
}

func (f singleLoggerLogFinder) Find(name string) Logger {
	return f.l
}
