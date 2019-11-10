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
		ctx := newContextWithID(r.Context())

		go func(ctx context.Context) {
			ctx = newContextWithID(ctx)
			time.Sleep(time.Second)
			logctx.From(ctx).Info("easy job done")
		}(ctx)

		go func(ctx context.Context) {
			ctx = newContextWithID(ctx)
			time.Sleep(10 * time.Second)
			logctx.From(ctx).Info("hard job done")
		}(ctx)
	})
}

var ctxIDCounter int64

func newContextWithID(ctx context.Context) context.Context {
	entry := logctx.From(ctx)
	prevCtxID, wasPrevCtxID := entry.Data["ctxID"].(int64)
	entry = entry.WithField("ctxID", atomic.AddInt64(&ctxIDCounter, 1))

	if wasPrevCtxID {
		entry.WithField("prevCtxID", prevCtxID).Info("new ctxID")
	} else {
		entry.Info("new ctxID")
	}

	return logctx.New(ctx, entry)
}
