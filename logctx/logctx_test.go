package logctx_test

import (
	"context"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"

	"github.com/go-logrusutil/logrusutil/logctx"

	log "github.com/sirupsen/logrus"
)

func Example_basic() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})

	// setting contextual log entry
	ctx := logctx.New(context.Background(), log.WithField("foo", "bar"))

	// retrieving context log entry, adding some data and emitting the log
	logctx.From(ctx).Info("hello world")
	// Output: level=info msg="hello world" foo=bar
}

func Example_hTTPRequestID() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})

	// contextual log middleware
	logMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := rand.Int()
			logEntry := log.WithField("request_id", reqID)
			logEntry.Info("request started")

			// setting contextual log entry
			ctx := logctx.New(r.Context(), logEntry)
			next.ServeHTTP(w, r.WithContext(ctx))

			logEntry.Info("request finished")
		})
	}

	// handler retrieving request contextual log entry, adding some data and emitting the log
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logctx.From(ctx).WithField("foo", "bar").Info("foobar created")
	})

	// run HTTP server with logging middleware
	ts := httptest.NewServer(logMiddleware(handler))
	defer ts.Close()

	// make request
	http.Get(ts.URL)
	// Output:
	// level=info msg="request started" request_id=5577006791947779410
	// level=info msg="foobar created" foo=bar request_id=5577006791947779410
	// level=info msg="request finished" request_id=5577006791947779410
}

func Example_goroutineID() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})

	// spawnGoroutine creates runs new goroutine with contextual log entries that has goroutine IDs
	// returns channel which closes when the goroutine
	// logs error and sends the panic value through the channel if goroutine panicked
	var goroutineIDCounter int64
	spawnGoroutine := func(ctx context.Context, fn func(context.Context)) <-chan interface{} {
		const LogFieldGoroutineID = "grtnID"
		const LogFieldGoroutineParentID = "grtnPrntID"

		entry := logctx.From(ctx)
		if gortnID, ok := entry.Data[LogFieldGoroutineID].(int64); ok {
			entry = entry.WithField(LogFieldGoroutineParentID, gortnID)
		}
		entry = entry.WithField(LogFieldGoroutineID, atomic.AddInt64(&goroutineIDCounter, 1))
		newCtx := logctx.New(ctx, entry)
		done := make(chan interface{})
		go func() {
			defer func() {
				if r := recover(); r != nil {
					entry.
						//	WithField("stack", string(debug.Stack())).
						WithField("panic", r).
						Error("goroutine panicked")
					done <- r
				}
				close(done)
			}()
			fn(newCtx)
		}()
		return done
	}

	// use spawnGoroutine and contextual logging
	<-spawnGoroutine(context.Background(), func(ctx context.Context) {
		logEntry := logctx.From(ctx).WithField("foo", "bar")
		logEntry.Info("first child goroutine started")

		<-spawnGoroutine(ctx, func(ctx context.Context) {
			panic("panic from second child")
		})

		logEntry.Info("first child goroutine finished")
	})
	// Output:
	// level=info msg="first child goroutine started" foo=bar grtnID=1
	// level=error msg="goroutine panicked" grtnID=2 grtnPrntID=1 panic="panic from second child"
	// level=info msg="first child goroutine finished" foo=bar grtnID=1
}

// IMPORTANT: this test is a the end because it alters global DefaultLogEntry,
// yet I want to have the examples simple

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
