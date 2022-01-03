package logx_test

import (
	"log"
	"os"

	"github.com/cmstar/go-logx"
)

func ExampleLoggerOp() {
	// StdLogger use os.Stderr by default, to make the example work,
	// we redirect the output to os.Stdout.
	l := logx.NewStdLogger(log.New(os.Stdout, "", 0))

	// Create the LoggerOp instance.
	op := logx.Op(l)

	// Use the shortcuts methods to write logs.
	op.Debug("a debug message")
	op.Infof("%d+%d=%d", 1, 2, 3)
	op.Errorkv("Name", "John", "Age", 55)

	// LoggerOp implements the Logger interface, so you can use the
	// raw methods of Logger.
	op.Log(logx.LevelWarn, "from Log() method", "key", "value")

	// Output:
	// DEBUG a debug message
	// INFO 1+2=3
	// ERROR  Name=John Age=55
	// WARN from Log() method key=value
}
