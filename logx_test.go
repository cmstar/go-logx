package logx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel_String(t *testing.T) {
	a := assert.New(t)
	a.Equal("DEBUG", LevelDebug.String())
	a.Equal("INFO", LevelInfo.String())
	a.Equal("WARN", LevelWarn.String())
	a.Equal("ERROR", LevelError.String())
	a.Equal("FATAL", LevelFatal.String())
	a.Equal("DEBUG|FATAL", (LevelDebug | LevelFatal).String())
}

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
	a.Equal("UNKNOWN", LevelToString(LevelFatal<<1))

	a.Equal("ERROR|FATAL", LevelToString(LevelBeyondError))
	a.Equal("WARN|ERROR|FATAL", LevelToString(LevelBeyondWarn))
	a.Equal("INFO|WARN|ERROR|FATAL", LevelToString(LevelBeyondInfo))
	a.Equal("DEBUG|INFO|WARN|ERROR|FATAL", LevelToString(LevelBeyondDebug))
}
