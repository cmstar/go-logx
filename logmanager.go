package logging

import "sync"

// DefaultManager is the globally shared LogManager.
var DefaultManager *LogManager = new(LogManager)

// LogManager is a simple implementation of LogFinder, which can be used to manage Loggers.
// It is safe for concurrent use.
type LogManager struct {
	loggers sync.Map
}

// NewManager creates a new instance of LogManager.
func NewManager() *LogManager {
	return &LogManager{}
}

var _ LogFinder = (*LogManager)(nil)

// Find returns the Logger instance with the specific name.
// If the name does not exist, returns nil.
func (m *LogManager) Find(name string) Logger {
	if l, ok := m.loggers.Load(name); ok {
		return l.(Logger)
	}
	return nil
}

// Set register a named logger to the current LogManager.
// If a logger with the name already exists, it will be replaced.
func (m *LogManager) Set(name string, logger Logger) {
	m.loggers.Store(name, logger)
}

// Unset removes a logger with the specified name from the current LogManager.
// If the name does not exists, the function is no-op.
func (m *LogManager) Unset(name string) {
	m.loggers.Delete(name)
}
