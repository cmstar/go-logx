package logx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogManager(t *testing.T) {
	m := NewManager()
	a := assert.New(t)

	checkLogger := func(name string, want Logger) {
		got := m.Find(name)
		a.Equal(want, got, "finding %v", name)
	}

	checkLogger("a", nil)
	m.Delete("") // No-op.

	type namedLogger struct {
		Logger
		name string
	}

	loggerA := namedLogger{name: "a"}
	m.Set("A", loggerA)

	loggerC1 := namedLogger{name: "a.b.c1"}
	m.Set("a.B.c1", loggerC1)

	loggerC2 := namedLogger{name: "a.b.c2"}
	m.Set("a.b.C2", loggerC2)

	// Loggers: root, a, a.b.c1, a.b.c2
	checkLogger("a", loggerA)
	checkLogger("a.B", loggerA)
	checkLogger("a.B.x", loggerA)
	checkLogger("a.B.x.y", loggerA)
	checkLogger("a.x", loggerA)
	checkLogger("a.x.y", loggerA)
	checkLogger("a.x.y.z", loggerA)

	// The heading dot is ignored.
	checkLogger(".A", loggerA)
	checkLogger(".a.B", loggerA)

	checkLogger("A.b.c1", loggerC1)
	checkLogger("a.B.c1.D", loggerC1)
	checkLogger("a.B.c1.D.e", loggerC1)

	checkLogger("a.B.C2", loggerC2)
	checkLogger("a.b.c2.e", loggerC2)

	checkLogger("", nil)
	checkLogger("a2", nil)
	checkLogger("b", nil)

	root := namedLogger{name: ""}
	m.Set("", root)
	checkLogger("", root)
	checkLogger("a2", root)
	checkLogger(".a", loggerA)
	checkLogger(".a.b", loggerA)

	m.Delete("a") // -> Loggers: root, a.b.c1, a.b.c2
	checkLogger("a", root)
	checkLogger("a.B", root)
	checkLogger("a.b.c1", loggerC1)
	checkLogger("a.b.c2", loggerC2)

	loggerB := namedLogger{name: "b"}
	m.Set("a.b", loggerB) // -> Loggers: root, a.b, a.b.c1, a.b.c2
	checkLogger("a", root)
	checkLogger("a.b", loggerB)
	checkLogger("a.b.c1", loggerC1)
	checkLogger("a.b.c2", loggerC2)

	m.Delete("a")      // No-op.
	m.Delete("a.b.x")  // No-op.
	m.Delete("a.b.c1") // -> Loggers: root, a.b, a.b.c2
	checkLogger("A.b.c1", loggerB)
	checkLogger("a.B.c1.D", loggerB)
	checkLogger("a.B.c1.D.e", loggerB)

	m.Delete("") // -> Loggers: a.b, a.b.c2
	checkLogger("a", nil)
	checkLogger("a.b", loggerB)
	checkLogger("a.b.x", loggerB)
	checkLogger("a.B.c2", loggerC2)
	checkLogger("a.b.c2.x", loggerC2)

	// Check the data structure.
	// Loggers: a.b, a.b.c2
	a.NotNil(m.nodes)
	a.Nil(m.nodes.logger)
	a.Equal(1, m.nodes.num)

	v, ok := m.nodes.children.Load("a")
	require.True(t, ok)
	nodeA := v.(*loggerNode)
	a.Nil(nodeA.logger)
	a.Equal(1, nodeA.num)

	v, ok = nodeA.children.Load("b")
	require.True(t, ok)
	nodeB := v.(*loggerNode)
	a.NotNil(nodeB.logger) // Logger 'a.b'.
	a.Equal(1, nodeB.num)

	v, ok = nodeB.children.Load("c2")
	require.True(t, ok)
	nodeC2 := v.(*loggerNode)
	a.NotNil(nodeC2.logger) // Logger 'a.b.c2'.
	a.Equal(0, nodeC2.num)

	// Remove all loggers.
	m.Delete("a.b.c2")
	a.Equal(0, nodeB.num)

	m.Delete("a.b")
	a.Equal(0, m.nodes.num)
}

func TestLogManager_Op(t *testing.T) {
	m := NewManager()
	assert.Nil(t, m.Op("x"))

	l := NewStdLogger(nil)
	m.Set("name", l)
	op := m.Op("name")
	assert.NotNil(t, op)
	assert.Equal(t, l, op.Logger)
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

	got = m.splitName(".a")
	assert.Equal(t, []string{"a"}, got)

	got = m.splitName(".")
	assert.Equal(t, []string{""}, got)

	got = m.splitName("..a..")
	assert.Equal(t, []string{"", "a", "", ""}, got)
}
