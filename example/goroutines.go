package main

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/go-logrusutil/logrusutil/logctx"
)

func registerGo(mux *http.ServeMux) {
	mux.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		ctx := contextWithNewJobID(r.Context())

		go func(ctx context.Context) {
			ctx = contextWithNewJobID(ctx)
			time.Sleep(time.Second)
			logctx.From(ctx).Info("easy job done")
		}(ctx)

		go func(ctx context.Context) {
			ctx = contextWithNewJobID(ctx)
			time.Sleep(10 * time.Second)
			logctx.From(ctx).Info("hard job done")
		}(ctx)
	})
}

var jobIDCounter int64

func contextWithNewJobID(ctx context.Context) context.Context {
	entry := logctx.From(ctx)
	prevJobID, wasPrevJobID := entry.Data["jobID"].(int64)
	entry = entry.WithField("jobID", atomic.AddInt64(&jobIDCounter, 1))

	if wasPrevJobID {
		entry.WithField("parentJobID", prevJobID).Info("new jobID")
	} else {
		entry.Info("new jobID")
	}

	return logctx.New(ctx, entry)
}
