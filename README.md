# logrusutil :hammer: [![GoDoc](https://godoc.org/github.com/logrusutil/v1?status.svg)](https://godoc.org/github.com/logrusutil/v1) ![Build Status](https://github.com/logrusutil/v1/workflows/build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/logrusutil/v1)](https://goreportcard.com/report/github.com/logrusutil/v1) [![golangci](https://golangci.com/badges/github.com/logrusutil/v1.svg)](https://golangci.com/r/github.com/logrusutil/v1)

Small, easy to use, yet powerful utility packages for <https://github.com/sirupsen/logrus>.

## `logctx` package [![GoDoc](https://godoc.org/github.com/logrusutil/v1/logctx?status.svg)](https://godoc.org/github.com/logrusutil/v1/logctx)

Add a log entry to the context using `logctx.New(ctx, logEntry)`. Retrieve the log entry using `logctx.From(ctx)`.

## `errfield` package [![GoDoc](https://godoc.org/github.com/logrusutil/v1/errfield?status.svg)](https://godoc.org/github.com/logrusutil/v1/errfield)

Wrap an error with fields using `errfield.Add(err, "key", value)`. Use `errfield.Formatter` to log the fields in a structured way.
