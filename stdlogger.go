package logx

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// StdLogger sends all log messages to the UnderlyingLogger which is a Logger of the standard library.
// If UnderlyingLogger is nil, log messages are sent to os.Stderr.
//
// Log messages are formatted with fmt.Sprintf(), the message format is:
//   LEVEL MESSAGE KEY1=VALUE1[ KEY2=VALUE2[ KYE3=VALUE3[...]]]
//
type StdLogger struct {
	// UnderlyingLogger receives formatted log messages.
	// If it is nil, os.Stderr will be used as the Logger.
	UnderlyingLogger *log.Logger

	mu sync.Mutex
}

// NewStdLogger creates a new StdLogger with the given underlyingLogger, which is used to receive
// log messages. If underlyingLogger is nil, log messages are sent to log.Default() which uses os.Stderr.
//
// The following code uses a underlying logger which sends messages to os.Stdin:
//   logger := logx.NewStdLogger(log.New(os.Stdin, "", log.LstdFlags))
//
func NewStdLogger(underlyingLogger *log.Logger) Logger {
	return &StdLogger{
		UnderlyingLogger: underlyingLogger,
	}
}

// Log implements Logger.Log().
func (logger *StdLogger) Log(level Level, message string, keyValues ...interface{}) error {
	logger.mu.Lock()
	defer logger.mu.Unlock()

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

// LogFn implements Logger.LogFn().
func (logger *StdLogger) LogFn(level Level, messageFactory func() (string, []interface{})) error {
	message, keyValues := messageFactory()
	return logger.Log(level, message, keyValues...)
}
