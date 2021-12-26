package logx

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	underlyingLogger := log.New(buf, "", 0)
	stdLogger := NewStdLogger(underlyingLogger)

	type args struct {
		level     Level
		message   string
		keyValues []interface{}
		want      string
	}

	tests := []struct {
		name string
		args args
	}{
		{"empty", args{
			level:     LevelDebug,
			message:   "",
			keyValues: []interface{}{},
			want:      "DEBUG \n",
		}},

		{"simple", args{
			level:     LevelInfo,
			message:   "simple",
			keyValues: []interface{}{},
			want:      "INFO simple\n",
		}},

		{"keyvalue", args{
			level:   LevelWarn,
			message: "msg",
			keyValues: []interface{}{
				"k1", 1,
				"k2", "v2",
			},
			want: "WARN msg k1=1 k2=v2\n",
		}},

		{"keyvalue-odd", args{
			level:   LevelWarn,
			message: "msg",
			keyValues: []interface{}{
				"k1", 1,
				"k2", "v2",
				"v3",
			},
			want: "WARN msg k1=1 k2=v2 UNKNOWN=v3\n",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()

			err := stdLogger.Log(tt.args.level, tt.args.message, tt.args.keyValues...)
			assert.NoError(t, err)

			msg := buf.String()
			assert.Equal(t, tt.args.want, msg)
		})
	}
}
