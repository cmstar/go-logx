package logxtest

import (
	"testing"

	"github.com/cmstar/go-logx"
	"github.com/stretchr/testify/assert"
)

func TestLogRecorder(t *testing.T) {
	r := NewRecorder()
	r.Log(logx.LevelDebug, "debug-msg", "k1", "v1")
	r.Log(logx.LevelWarn, "warn-msg", "k2", "v2")
	r.LogFn(logx.LevelFatal, func() (string, []interface{}) { return "fatal-msg", []interface{}{"k3", "v3"} })

	a := assert.New(t)
	a.Equal(3, len(r.Messages))

	lines := r.Lines()
	a.Equal(3, len(lines))
	a.Equal("DEBUG debug-msg k1=v1\n", lines[0])
	a.Equal("WARN warn-msg k2=v2\n", lines[1])
	a.Equal("FATAL fatal-msg k3=v3\n", lines[2])

	wholeLog := `DEBUG debug-msg k1=v1
WARN warn-msg k2=v2
FATAL fatal-msg k3=v3
`
	a.Equal(wholeLog, r.String())
}
