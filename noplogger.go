package logging

// NopLogger is a no-op logger, the Log function do nothing and returns no error.
// It is safe for concurrent use.
var NopLogger Logger = new(nopLogger)

type nopLogger struct{}

func (logger *nopLogger) Log(level Level, message string, keyValues ...interface{}) error {
	return nil
}

func (logger *nopLogger) LogFn(level Level, messageFactory func() (string, []interface{})) error {
	return nil
}
