package logx

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerOp(t *testing.T) {
	buf := new(bytes.Buffer)
	underlyingLogger := log.New(buf, "", 0)
	op := Op(NewStdLogger(underlyingLogger))

	op.Debug("Debug msg")
	op.Debugf("Debug %v", 1)
	op.Debugkv("k1", 11, "k2", 12)

	op.Info("Info msg")
	op.Infof("Info %v", 2)
	op.Infokv("k1", 21, "k2", 22)

	op.Warn("Warn msg")
	op.Warnf("Warn %v", 3)
	op.Warnkv("k1", 31, "k2", 32)

	op.Error("Error msg")
	op.Errorf("Error %v", 4)
	op.Errorkv("k1", 41, "k2", 42)

	op.Fatal("Fatal msg")
	op.Fatalf("Fatal %v", 5)
	op.Fatalkv("k1", 51, "k2", 52)

	got := buf.String()
	want := `DEBUG Debug msg
DEBUG Debug 1
DEBUG  k1=11 k2=12
INFO Info msg
INFO Info 2
INFO  k1=21 k2=22
WARN Warn msg
WARN Warn 3
WARN  k1=31 k2=32
ERROR Error msg
ERROR Error 4
ERROR  k1=41 k2=42
FATAL Fatal msg
FATAL Fatal 5
FATAL  k1=51 k2=52
`
	assert.Equal(t, want, got)
}
