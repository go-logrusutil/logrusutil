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
			// Output: time="2019-11-11T20:50:30.9971817+01:00" level=info msg="easy job done" app=example jobID=2 reqID=1443635317331776148
		}(ctx)

		go func(ctx context.Context) {
			ctx = contextWithNewJobID(ctx)
			time.Sleep(10 * time.Second)
			logctx.From(ctx).Info("hard job done")
			// Output: time="2019-11-11T20:50:40.029896+01:00" level=info msg="hard job done" app=example jobID=3 reqID=1443635317331776148
		}(ctx)
	})
}

var jobIDCounter int64

// contextWithNewJobID creates new context with next jobID and adds it to the contextual log entry
func contextWithNewJobID(ctx context.Context) context.Context {
	entry := logctx.From(ctx)
	prevJobID, wasPrevJobID := entry.Data["jobID"].(int64)
	entry = entry.WithField("jobID", atomic.AddInt64(&jobIDCounter, 1))

	if wasPrevJobID {
		entry.WithField("parentJobID", prevJobID).Info("new jobID")
		// Output: time="2019-11-11T20:50:29.9299289+01:00" level=info msg="new jobID" app=example jobID=2 parentJobID=1 reqID=1443635317331776148
	} else {
		entry.Info("new jobID")
		// Output: time="2019-11-11T20:50:29.9289259+01:00" level=info msg="new jobID" app=example jobID=1 reqID=1443635317331776148
	}

	return logctx.New(ctx, entry)
}
