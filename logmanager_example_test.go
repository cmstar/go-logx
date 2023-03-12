package logx_test

import (
	"fmt"

	"github.com/cmstar/go-logx"
)

func ExampleLogManager() {
	// Create a LogManager, or you can use logx.DefaultManager directly.
	manager := logx.NewManager()

	// Define some custom Loggers.
	type root struct{ logx.Logger }
	type loggerA struct{ logx.Logger }
	type loggerC struct{ logx.Logger }
	type loggerE struct{ logx.Logger }

	// Register these Loggers.
	manager.Set("", root{logx.NopLogger})
	manager.Set("a", loggerA{logx.NopLogger})
	manager.Set("a.b.c", loggerC{logx.NopLogger})
	manager.Set("a.b.c.d.e", loggerE{logx.NopLogger})

	show := func(name string) {
		logger := manager.Find(name)
		if name == "" {
			name = "''"
		}
		fmt.Printf("%-10v: %T\n", name, logger)
	}

	/*
		`LogManager` uses case-insensitive header matching when finding Loggers.
		A name will be split by the dot(.) into several segments,
		when finding a name like 'A.B.C.D', `LogManager` finds the `Logger` in this order,
		returns the first found `Logger`:
		  - a.b.c.d
		  - a.b.c
		  - a.b
		  - a
		  - "" (empty string)
	*/
	show("")          // Found the root logger.
	show("A")         // Found logger 'a' (case-insensitive).
	show("a.B")       // 'a.b' not found, try finding 'a'.
	show("A.b.C")     // Found logger 'a.b.c'.
	show("a.b.c.d")   // 'a.b.c.d' not found, try finding 'a.b.c'.
	show("a.b.c.d.e") // 'a.b.c.d.e' not found, try finding 'a.b.c.d', then 'a.b.c'.
	show("x")         // 'x' not found, try finding ''.
	show("x.y")       // 'x.y' not found, try finding 'x', then ''.
	show("a.x")       // 'a.x' not found, try finding 'a'.
	show("a.x.y")     // 'a.x.y' not found, try finding 'a.x', then 'a'.

	// Delete a Logger.
	manager.Delete("a.b.c")
	fmt.Println("loggerC is deleted")
	show("a.b.c") // Fallback to logger 'a'.
	manager.Delete("")
	fmt.Println("the root logger is deleted")
	show("x") // Got nil because there's no root logger now.

	// Output:
	// ''        : logx_test.root
	// A         : logx_test.loggerA
	// a.B       : logx_test.loggerA
	// A.b.C     : logx_test.loggerC
	// a.b.c.d   : logx_test.loggerC
	// a.b.c.d.e : logx_test.loggerE
	// x         : logx_test.root
	// x.y       : logx_test.root
	// a.x       : logx_test.loggerA
	// a.x.y     : logx_test.loggerA
	// loggerC is deleted
	// a.b.c     : logx_test.loggerA
	// the root logger is deleted
	// x         : <nil>
}
