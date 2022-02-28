package logx_test

import (
	"testing"

	"github.com/cmstar/go-logx"
	"github.com/cmstar/go-logx/logxtest"
	"github.com/stretchr/testify/assert"
)

func TestFilterLevel(t *testing.T) {
	r := logxtest.NewRecorder()
	l := logx.FilterLevel(r, logx.LevelDebug|logx.LevelError)

	l.Log(logx.LevelDebug, "d")
	l.Log(logx.LevelInfo, "i")
	l.Log(logx.LevelWarn, "w")
	l.Log(logx.LevelError, "e")
	l.Log(logx.LevelFatal, "f")
	l.LogFn(logx.LevelDebug, func() (message string, keyValues []interface{}) { return "d2", nil })
	l.LogFn(logx.LevelInfo, func() (message string, keyValues []interface{}) { return "i2", nil })

	assert.Equal(t, 3, len(r.Messages))

	assert.Equal(t, "d", r.Messages[0].Message)
	assert.Equal(t, logx.LevelDebug, r.Messages[0].Level)

	assert.Equal(t, "e", r.Messages[1].Message)
	assert.Equal(t, logx.LevelError, r.Messages[1].Level)

	assert.Equal(t, "d2", r.Messages[2].Message)
	assert.Equal(t, logx.LevelDebug, r.Messages[2].Level)
}
