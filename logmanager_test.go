package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogManager(t *testing.T) {
	m := NewManager()

	checkLogger := func(name string, want Logger) {
		got := m.Find(name)
		assert.Equal(t, want, got, "finding %v", name)
	}

	checkLogger("a", nil)
	m.Delete("") // No-op.

	logger1 := NewStdLogger(nil)
	m.Set("A", logger1)

	logger2 := NewStdLogger(nil)
	m.Set("a.B.c", logger2)

	logger3 := NewStdLogger(nil)
	m.Set("a.b.D", logger3)

	checkLogger("a", logger1)
	checkLogger("a.B", logger1)
	checkLogger("a.B.x", logger1)
	checkLogger("a.B.x.y", logger1)
	checkLogger("a.x", logger1)
	checkLogger("a.x.y", logger1)
	checkLogger("a.x.y.z", logger1)

	checkLogger("A.b.c", logger2)
	checkLogger("a.B.c.D", logger2)
	checkLogger("a.B.c.D.e", logger2)

	checkLogger("a.B.d", logger3)
	checkLogger("a.b.d.e", logger3)

	checkLogger("", nil)
	checkLogger("a2", nil)
	checkLogger(".a", nil)
	checkLogger("b", nil)

	logger0 := NewStdLogger(nil)
	m.Set("", logger0)
	checkLogger("", logger0)
	checkLogger(".a", logger0)
	checkLogger(".a.b", logger0)
	checkLogger("a2", nil)

	m.Delete("a.b.c")
	checkLogger("A.b.c", logger1)
	checkLogger("a.B.c.D", logger1)
	checkLogger("a.B.c.D.e", logger1)

	m.Delete("a")
	checkLogger("a", nil)
	checkLogger("a.B", nil)

	// Check the data structure.
	a := assert.New(t)
	a.NotNil(m.nodes)
	a.Equal(2, len(m.nodes.children))
	nodeA, ok := m.nodes.children["a"]
	a.True(ok)
	a.Nil(nodeA.logger)
	a.Equal(1, len(nodeA.children))
	nodeB, ok := nodeA.children["b"]
	a.True(ok)
	a.Nil(nodeB.logger)
	a.Equal(1, len(nodeB.children))
	nodeD, ok := nodeB.children["d"]
	a.True(ok)
	a.Equal(0, len(nodeD.children))

	// Make sure other loggers are there.
	checkLogger("a.B.d", logger3)
	checkLogger("a.b.d.e", logger3)

	// Remove all logger.
	m.Delete("")
	m.Delete("a.b.d")
	a.Equal(0, len(m.nodes.children))
}

func TestLogManager_splitName(t *testing.T) {
	var m *LogManager

	got := m.splitName("")
	assert.Equal(t, []string{""}, got)

	got = m.splitName(" ")
	assert.Equal(t, []string{" "}, got)

	got = m.splitName("A")
	assert.Equal(t, []string{"a"}, got)

	got = m.splitName("a, b")
	assert.Equal(t, []string{"a, b"}, got)

	got = m.splitName("a.Bb.cC.dd")
	assert.Equal(t, []string{"a", "bb", "cc", "dd"}, got)
}
