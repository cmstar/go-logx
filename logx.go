// Package logx provides abstraction and some simple implementation of logging.
package logx

import (
	"fmt"
	"strings"
)

// Level defines the log level.
// If level A is lower than B, the integer value of A should be less than B.
type Level int8

// Levels can be a bitmap mask.
const (
	LevelDebug Level = 1 << iota // LevelDebug is the debug level.
	LevelInfo                    // LevelInfo is the info level.
	LevelWarn                    // LevelWarn is the warn level.
	LevelError                   // LevelError is the error level.
	LevelFatal                   // LevelFatal is the fatal level.
)

var _ fmt.Stringer = Level(0)

// String returns the string representation of Level.
// It's equivalent to LevelToString(v).
func (v Level) String() string {
	return LevelToString(v)
}

// Combined levels.
const (
	// LevelBeyondError combines LevelFatal, LevelError.
	LevelBeyondError = LevelError | LevelFatal

	// LevelBeyondError combines LevelFatal, LevelError, LevelWarn.
	LevelBeyondWarn = LevelWarn | LevelBeyondError

	// LevelBeyondError combines LevelFatal, LevelError, LevelWarn, LevelInfo.
	LevelBeyondInfo = LevelInfo | LevelBeyondWarn

	// LevelBeyondError combines LevelFatal, LevelError, LevelWarn, LevelInfo, LevelDebug.
	LevelBeyondDebug = LevelDebug | LevelBeyondInfo
)

// ParseLevel parses the given string to the corresponding Level.
// Returns -1 if the value cannot be parsed.
func ParseLevel(v string) Level {
	v = strings.ToUpper(v)
	switch v {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	case "FATAL":
		return LevelFatal
	}
	return -1
}

// LevelToString returns the string representation of Level.
//
// The string is in uppercase like DEBUG, INFO, WARN, ERROR, FATAL.
// Combined levels are split by '|', e.g. DEBUG|INFO|ERROR .
//
// If the given level is not defined, returns UNKNOWN.
func LevelToString(lv Level) string {
	res := ""
	a := func(s string) {
		if len(res) > 0 {
			res += "|"
		}
		res += s
	}

	if (lv & LevelDebug) == LevelDebug {
		a("DEBUG")
	}

	if (lv & LevelInfo) == LevelInfo {
		a("INFO")
	}

	if (lv & LevelWarn) == LevelWarn {
		a("WARN")
	}

	if (lv & LevelError) == LevelError {
		a("ERROR")
	}

	if (lv & LevelFatal) == LevelFatal {
		a("FATAL")
	}

	if len(res) == 0 {
		return "UNKNOWN"
	}
	return res
}

// Logger defines the logging action.
// All methods should be safe for concurrent use.
type Logger interface {
	// Log creates a log message at the given Level, returns an error if the creation failed.
	// The log message will not be processed if the given level is not enabled on the current logger.
	// If the generation of log messages is expensive, use LogFn instead.
	//
	// A log message can contain couples of key-value pairs to extend the message.
	// The elements at even indexes of keyValues are the keys; the odd indexes are the values.
	//
	// If the keys and values are unpaired, the last element should be treated as a value with key 'UNKNOWN'.
	//
	Log(level Level, message string, keyValues ...interface{}) error

	// LogFn is similar to the Log function, but generates the log message with a function.
	// It is useful when the generation of log messages is expensive.
	//
	// If messageFactory panics, it will not be handled.
	//
	LogFn(level Level, messageFactory func() (message string, keyValues []interface{})) error
}
