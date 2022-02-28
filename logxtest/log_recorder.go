// Package logxtest provides some types for testing the logx package.
package logxtest

import (
	"log"
	"strings"

	"github.com/cmstar/go-logx"
)

// NewRecorder returns a new instance of LogRecorder.
func NewRecorder() *LogRecorder {
	return new(LogRecorder)
}

// LogMessage is a log message recorded by LogRecorder.
type LogMessage struct {
	Level     logx.Level
	Message   string
	KeyValues []interface{}
}

// LogRecorder is an implementation of logx.Logger, that records log messages for test.
type LogRecorder struct {
	Messages []LogMessage // The recorded messages.
}

var _ logx.Logger = (*LogRecorder)(nil)

func (r *LogRecorder) Log(level logx.Level, message string, keyValues ...interface{}) error {
	r.Messages = append(r.Messages, LogMessage{
		Level:     level,
		Message:   message,
		KeyValues: keyValues,
	})
	return nil
}

func (r *LogRecorder) LogFn(level logx.Level, messageFactory func() (message string, keyValues []interface{})) error {
	m, k := messageFactory()
	return r.Log(level, m, k...)
}

// Lines returns a slice of strings, each element is a formatted log message.
// It formats log messages in the same manner of logx.StdLogger.
func (r *LogRecorder) Lines() []string {
	// We use StdLogger directly.
	buf := new(strings.Builder)
	stdLogger := logx.NewStdLogger(log.New(buf, "", 0))
	lines := make([]string, 0, len(r.Messages))
	for _, msg := range r.Messages {
		stdLogger.Log(msg.Level, msg.Message, msg.KeyValues...)
		line := buf.String()
		lines = append(lines, line)
		buf.Reset()
	}
	return lines
}

// String joins Lines() and returns the whole log as a string.
func (r *LogRecorder) String() string {
	buf := new(strings.Builder)
	stdLogger := logx.NewStdLogger(log.New(buf, "", 0))
	for _, msg := range r.Messages {
		stdLogger.Log(msg.Level, msg.Message, msg.KeyValues...)
	}
	return buf.String()
}
