package logx

import (
	"strings"
	"sync"
)

// DefaultManager is the globally shared LogManager.
var DefaultManager *LogManager = NewManager()

// LogManager is a simple implementation of LogFinder, which can be used to manage Loggers.
// It is safe for concurrent use.
//
// LogManager uses case-insensitive header matching when finding Loggers.
// A name will be splitted by the dot(.) into several segments, when finding a name like 'A.B.C.D',
// LogManager finds the Logger in this order, returns the first found `Logger`:
//   - a.b.c.d
//   - a.b.c
//   - a.b
//   - a
//   - "" (empty string)
//
// If no logger can be found, LogManger.Find() returns nil.
//
// The logger with the empty name is the root logger. The empty string can be treated as the
// first segment of all other names, e.g. The name 'a.b' is equivalent to '.a.b'.
//
// An empty string is a legal segment, that is, a logger name can be '.A..b',
// which will be splitted into ['', 'a', '', 'b'].
//
type LogManager struct {
	mu    sync.RWMutex // Manages locks.
	nodes *loggerNode  // The root node of the tree, whose logger field is always nil.
}

// loggerNode is a node in the tree that stores Loggers.
// Each node stores a segment of a logger name. The root node's segment field is always the empty string.
//
// e.g., there are Loggers with name 'a.b', '.a.d', 'a.b.c', '.x.y', '..h', they will be stored in the tree like:
//   root{ segment: '', logger: nil }
//     |- node{ segment: 'a', logger: nil }
//     |    |- node{ segment: 'b', logger: of('a.b') }
//     |    |    |- node{ segment: 'c', logger: of('a.b.c') }
//     |    |
//     |    |- node{ segment: 'd', logger: of('a.d') }
//     |
//     |- node{ segment: 'x', logger: nil }
//     |    |- node{ segment: 'y', logger: of('x.y') }
//     |
//     |- node{ segment: '', logger: nil }
//          |- node{ segment: 'h', logger: of('.h') }
//
// Note: '.a.d' is equivalent to 'a.d'; '.x.y' is equivalent to 'x.y'; '..h' is equivalent to '.h'.
//
type loggerNode struct {
	logger   Logger      // nil if this segment keeps no logger directly, thus loggers are kept on the children field.
	segment  string      // The segment of the node.
	parent   *loggerNode // Points to the parent node, nil if the current node is root.
	children map[string]*loggerNode
}

// NewManager creates a new instance of LogManager.
func NewManager() *LogManager {
	return &LogManager{}
}

var _ LogFinder = (*LogManager)(nil)

// Find returns the Logger instance with the specific name.
// If the name cannot be found, returns nil.
func (m *LogManager) Find(name string) Logger {
	m.mu.RLock()
	defer func() { m.mu.RUnlock() }()

	_, lastNonNil := m.doFind(name)
	if lastNonNil == nil {
		return nil
	}
	return lastNonNil.logger
}

// Set registers a named logger to the current LogManager.
// If a logger with the name already exists, it will be replaced.
func (m *LogManager) Set(name string, logger Logger) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Try to initialize the root node.
	if m.nodes == nil {
		m.nodes = new(loggerNode)
	}

	if name == "" {
		m.nodes.logger = logger
		return
	}

	var current, next *loggerNode
	var hasChild bool
	segments := m.splitName(name)
	current = m.nodes

	for i := 0; i < len(segments); i++ {
		seg := segments[i]

		if current.children == nil {
			hasChild = false
		} else {
			next, hasChild = current.children[seg]
		}

		if !hasChild {
			if current.children == nil {
				current.children = make(map[string]*loggerNode)
			}
			next = &loggerNode{
				segment: seg,
				parent:  current,
			}
			current.children[seg] = next
		}

		current = next
	}

	current.logger = logger
}

// Delete removes a logger with the specified name from the current LogManager.
// If the name does not exists, the function is no-op.
func (m *LogManager) Delete(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	current, _ := m.doFind(name)
	if current == nil {
		return
	}

	current.logger = nil

	// Try to purge the parent node if the node keeps no logger now.
	// This action can be performed recursively.
	for {
		if current.parent == nil || len(current.children) > 0 {
			break
		}

		delete(current.parent.children, current.segment)
		current = current.parent
	}
}

// doFind performs finding on the logger tree.
// @current is the node which keeps the logger of the given name, nil if @name cannot be located.
// @lastNonNil is the nearest ancestor node of @current whose logger field is not nil.
func (m *LogManager) doFind(name string) (current, lastNonNil *loggerNode) {
	if m.nodes == nil {
		return nil, nil
	}

	if name == "" {
		return m.nodes, m.nodes
	}

	var next *loggerNode
	var hasChild bool
	var i int
	segments := m.splitName(name)
	current = m.nodes
	lastNonNil = current

	for i = 0; i < len(segments); i++ {
		seg := segments[i]

		if current.children == nil {
			break
		} else {
			next, hasChild = current.children[seg]
		}

		if !hasChild {
			break
		}

		current = next

		if current.logger != nil {
			lastNonNil = current
		}
	}

	return
}

// splitName splits the given name by a dot(.) into a group of lowercase segments.
// If the first segment is a empty string, it will be ignored.
func (*LogManager) splitName(name string) []string {
	name = strings.ToLower(name)
	segments := strings.Split(name, ".")
	if len(segments) > 1 && segments[0] == "" {
		return segments[1:]
	}
	return segments
}
