package logging

import (
	"fmt"
	"log"
	"strings"
)

// StdLogger sends all log messages to the UnderlyingLogger which is a Logger of the standard library.
// If UnderlyingLogger is nil, log messages are sent to log.Default().
//
// Log messages are formatted with fmt.Sprintf(), the message format is:
//   LEVEL MESSAGE KEY1=VALUE1[ KEY2=VALUE2[ KYE3=VALUE3[...]]]
//
type StdLogger struct {
	// UnderlyingLogger receives formatted log messages.
	// If it is nil, log.Default() will be used.
	UnderlyingLogger *log.Logger
}

// NewStdLogger creates a new StdLogger with the given underlyingLogger, which is used to receive
// log messages. If underlyingLogger is nil, log messages are sent to log.Default().
func NewStdLogger(underlyingLogger *log.Logger) Logger {
	return &StdLogger{underlyingLogger}
}

func (logger *StdLogger) Log(level Level, message string, keyValues ...interface{}) error {
	builder := new(strings.Builder)

	// Format: LEVEL MESSAGE KEY=VALUE KEY=VALUE
	builder.WriteString(LevelToString(level))
	builder.WriteByte(' ')
	builder.WriteString(message)

	length := len(keyValues)
	if length > 0 {
		for i := 0; i < length-1; i += 2 {
			segment := fmt.Sprintf(" %v=%v", keyValues[i], keyValues[i+1])
			builder.WriteString(segment)
		}

		if length%2 != 0 {
			segment := fmt.Sprintf(" UNKNOWN=%v", keyValues[length-1])
			builder.WriteString(segment)
		}
	}

	underlyingLogger := logger.UnderlyingLogger
	if underlyingLogger == nil {
		underlyingLogger = log.Default()
	}

	line := builder.String()
	underlyingLogger.Println(line)
	return nil
}

func (logger *StdLogger) LogFn(level Level, messageFactory func() (string, []interface{})) error {
	message, keyValues := messageFactory()
	return logger.Log(level, message, keyValues...)
}
