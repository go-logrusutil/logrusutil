package logctx_test

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-logrusutil/logrusutil/logctx"
)

func Example() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})

	// setting contextual log entry
	ctx := logctx.New(context.Background(), log.WithField("foo", "bar"))

	// retrieving context log entry, adding some data and emitting the log
	logctx.From(ctx).Info("hello world")
	// Output: level=info msg="hello world" foo=bar
}

// IMPORTANT: this test is a the end because it alters global DefaultLogEntry,
// yet I want to have the example as simple as possible

//
func Example_defaultLogEntry() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})

	// set the default log entry
	logctx.DefaultLogEntry = log.WithField("foo", "bar")

	// get a log entry from context for which a contextual entry was not set
	logctx.From(context.Background()).Info("hello world")
	// Output: level=info msg="hello world" foo=bar
}
