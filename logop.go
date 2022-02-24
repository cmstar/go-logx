package logx

import "fmt"

// LoggerOp wraps a Logger interface, provides a group of shortcut operations
// to call the methods of the Logger.
type LoggerOp struct {
	Logger
}

// Op returns a LoggerOp which wraps the given Logger.
// If the given logger is nil, the LoggerOp will use a NopLogger.
func Op(logger Logger) *LoggerOp {
	if logger == nil {
		logger = NopLogger
	}
	return &LoggerOp{logger}
}

// Debug calls Logger.Log() using LevelDebug, without the key-value part.
func (op *LoggerOp) Debug(msg string) {
	op.Log(LevelDebug, msg)
}

// Debugf calls Logger.Log() using LevelDebug, without the key-value part,
// and formats the message with fmt.Sprintf().
func (op *LoggerOp) Debugf(format string, args ...interface{}) {
	op.Log(LevelDebug, fmt.Sprintf(format, args...))
}

// Debugkv calls Logger.Log() using LevelDebug, without the message part.
func (op *LoggerOp) Debugkv(keyValues ...interface{}) {
	op.Log(LevelDebug, "", keyValues...)
}

// Info calls Logger.Log() using LevelInfo, without the key-value part.
func (op *LoggerOp) Info(msg string) {
	op.Log(LevelInfo, msg)
}

// Infof calls Logger.Log() using LevelInfo, without the key-value part,
// and formats the message with fmt.Sprintf().
func (op *LoggerOp) Infof(format string, args ...interface{}) {
	op.Log(LevelInfo, fmt.Sprintf(format, args...))
}

// Infokv calls Logger.Log() using LevelInfo, without the message part.
func (op *LoggerOp) Infokv(keyValues ...interface{}) {
	op.Log(LevelInfo, "", keyValues...)
}

// Warn calls Logger.Log() using LevelWarn, without the key-value part.
func (op *LoggerOp) Warn(msg string) {
	op.Log(LevelWarn, msg)
}

// Warnf calls Logger.Log() using LevelWarn, without the key-value part,
// and formats the message with fmt.Sprintf().
func (op *LoggerOp) Warnf(format string, args ...interface{}) {
	op.Log(LevelWarn, fmt.Sprintf(format, args...))
}

// Warnkv calls Logger.Log() using LevelWarn, without the message part.
func (op *LoggerOp) Warnkv(keyValues ...interface{}) {
	op.Log(LevelWarn, "", keyValues...)
}

// Error calls Logger.Log() using LevelError, without the key-value part.
func (op *LoggerOp) Error(msg string) {
	op.Log(LevelError, msg)
}

// Errorf calls Logger.Log() using LevelError, without the key-value part,
// and formats the message with fmt.Sprintf().
func (op *LoggerOp) Errorf(format string, args ...interface{}) {
	op.Log(LevelError, fmt.Sprintf(format, args...))
}

// Errorkv calls Logger.Log() using LevelError, without the message part.
func (op *LoggerOp) Errorkv(keyValues ...interface{}) {
	op.Log(LevelError, "", keyValues...)
}

// Fatal calls Logger.Log() using LevelFatal, without the key-value part.
func (op *LoggerOp) Fatal(msg string) {
	op.Log(LevelFatal, msg)
}

// Fatalf calls Logger.Log() using LevelFatal, without the key-value part,
// and formats the message with fmt.Sprintf().
func (op *LoggerOp) Fatalf(format string, args ...interface{}) {
	op.Log(LevelFatal, fmt.Sprintf(format, args...))
}

// Fatalkv calls Logger.Log() using LevelFatal, without the message part.
func (op *LoggerOp) Fatalkv(keyValues ...interface{}) {
	op.Log(LevelFatal, "", keyValues...)
}
