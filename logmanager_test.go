package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogManager(t *testing.T) {
	const name = "name"
	m := NewManager()

	got := m.GetLogger(name)
	assert.Nil(t, got)

	m.Set(name, NopLogger)
	got = m.GetLogger(name)
	assert.Equal(t, NopLogger, got)

	// Replacing.
	l := NewStdLogger(nil)
	m.Set(name, l)
	got = m.GetLogger(name)
	assert.Equal(t, got, l)

	m.Unset(name)
	got = m.GetLogger(name)
	assert.Nil(t, got)

	m.Unset(name) // No-op.
}
