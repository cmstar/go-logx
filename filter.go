package logx

// FilterLevel wraps the given Logger, returns a new Logger which can filter log messages by the Level mask.
// If the level of a log message is included in the mask, it is logged; otherwise it is dropped.
//
// e.g. If levelMast is Warn|Error|Fatal , only messages with level greater than Warn are logged.
//
func FilterLevel(raw Logger, levelMask Level) Logger {
	return logLevelFilter{raw, levelMask}
}

// logLevelFilter is a Logger which can filter log messages by the log level.
type logLevelFilter struct {
	logger    Logger
	levelMask Level
}

func (f logLevelFilter) Log(level Level, message string, keyValues ...interface{}) error {
	if (f.levelMask & level) != level {
		return nil
	}
	return f.logger.Log(level, message, keyValues...)
}

func (f logLevelFilter) LogFn(level Level, messageFactory func() (message string, keyValues []interface{})) error {
	if (f.levelMask & level) != level {
		return nil
	}
	return f.logger.LogFn(level, messageFactory)
}
