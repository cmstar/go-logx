# logx - Simple logging abstraction

[![GoDoc](https://pkg.go.dev/badge/github.com/cmstar/go-logx)](https://pkg.go.dev/github.com/cmstar/go-logx)
[![Go](https://github.com/cmstar/go-logx/workflows/Go/badge.svg)](https://github.com/cmstar/go-logx/actions?query=workflow%3AGo)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)

There are so many logging frameworks, it's not necessary to write a new one.

This package only provides some logging abstraction. The purpose is to be used as a middleware between business codes and other logging frameworks.

## Logger

The `Logger` interface has two methods:
- `Log()`: writes a log message.
- `LogFn()`: similar to `Log`, but accept a factory function to produce the log message.

There are two built-in loggers:
- `NopLogger`: an empty logger that logs nothing.
- `StdLogger`: writes logs to the standard output.

For more details, see the [GoDoc](https://pkg.go.dev/github.com/cmstar/go-logx#Logger).

## LogManager

`LogManager` is used for managing `Logger`s. It's safe for concurrent use.

Methods:
- `Set(name, Logger)`: register a logger with the given name into the `LogManager` instance.
- `Find(name)`: get a logger with the given name.
- `Delete(name)`: deleted a registered logger from the instance.

`LogManager` uses case-insensitive header matching when finding Loggers. A name will be splitted by the dot(.) into several segments, when finding a name like 'A.B.C.D', `LogManager` finds the `Logger` in this order, returns the first found `Logger`:
- a.b.c.d
- a.b.c
- a.b
- a
- "" (empty string)

If no logger can be found, `LogManger.Find()` returns `nil`.

The logger with the empty name is the root logger. The empty string can be treated as the first segment of all other names, e.g. The name 'a.b' is equivalent to '.a.b'.

> It's similar to the `LoaManager` class in `log4j` from `Java`/`Common.Logging` from `.net`

For more details, see the [Example](https://pkg.go.dev/github.com/cmstar/go-logx#example-LogManager).
