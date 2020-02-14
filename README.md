# logrusutil :hammer: [![GoDoc](https://godoc.org/github.com/go-logrusutil/logrusutil?status.svg)](https://godoc.org/github.com/go-logrusutil/logrusutil) ![Build Status](https://github.com/go-logrusutil/logrusutil/workflows/build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/go-logrusutil/logrusutil)](https://goreportcard.com/report/github.com/go-logrusutil/logrusutil) [![golangci](https://golangci.com/badges/github.com/go-logrusutil/logrusutil.svg)](https://golangci.com/r/github.com/go-logrusutil/logrusutil)

Small, easy to use, yet powerful utility packages for [**logrus**](https://github.com/sirupsen/logrus).

## `logctx` package [![GoDoc](https://godoc.org/github.com/go-logrusutil/logrusutil/logctx?status.svg)](https://godoc.org/github.com/go-logrusutil/logrusutil/logctx)

Add a log entry to the context using `logctx.New(ctx, logEntry)` or simply add a new log field using `logctx.AddField(ctx, key, value)`. Retrieve the log entry using `logctx.From(ctx)`.

[Read more.](https://pajak.home.blog/2019/12/12/logging-6-correlation-log-field/)

## `errfield` package [![GoDoc](https://godoc.org/github.com/go-logrusutil/logrusutil/errfield?status.svg)](https://godoc.org/github.com/go-logrusutil/logrusutil/errfield)

Wrap an error with fields using `errfield.Add(err, "key", value)`. Use `errfield.Formatter` to log the fields in a structured way.

[Read more.](https://pajak.home.blog/2019/12/15/logging-7-contextual-error-data-log-field/)

## [Examples](example)

Check the [**sample HTTP server application**](example) and its [**Step by step guide**](example/README.md#step-by-step-guide), to see **logrusutil** in action.
