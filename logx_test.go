package logx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	a := assert.New(t)
	a.Equal(LevelDebug, ParseLevel("debug"))
	a.Equal(LevelInfo, ParseLevel("INFO"))
	a.Equal(LevelWarn, ParseLevel("warn"))
	a.Equal(LevelError, ParseLevel("Error"))
	a.Equal(LevelFatal, ParseLevel("fatal"))
	a.Equal(Level(-1), ParseLevel("x"))
}

func TestLevelToString(t *testing.T) {
	a := assert.New(t)
	a.Equal("DEBUG", LevelToString(LevelDebug))
	a.Equal("INFO", LevelToString(LevelInfo))
	a.Equal("WARN", LevelToString(LevelWarn))
	a.Equal("ERROR", LevelToString(LevelError))
	a.Equal("FATAL", LevelToString(LevelFatal))
	a.Equal("UNKNOWN", LevelToString(Level(100)))
}
