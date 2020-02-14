// Package logctx adds contextual logging to logrus.
package logctx

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Default is used to create a new log entry if there is none in the context.
var Default = logrus.NewEntry(logrus.StandardLogger())

type contextKey struct{}

// New creates a new context with log entry.
func New(ctx context.Context, logEntry *logrus.Entry) context.Context {
	return context.WithValue(ctx, contextKey{}, logEntry)
}

// From returns the log entry from the context.
// Returns log entry from Default if there is no log entry in the context.
func From(ctx context.Context) *logrus.Entry {
	if entry, ok := ctx.Value(contextKey{}).(*logrus.Entry); ok {
		return entry
	}
	return Default
}

// AddField adds a log field to the contexual log entry.
func AddField(ctx context.Context, key string, value interface{}) context.Context {
	entry := From(ctx).WithField(key, value)
	return New(ctx, entry)
}
