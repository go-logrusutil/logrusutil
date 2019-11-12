package main

import (
	"math/rand"
	"net/http"

	"github.com/go-logrusutil/logrusutil/logctx"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add a random reqID field for each request
		reqID := rand.Int()
		logEntry := logctx.Default.WithField("reqID", reqID)
		logEntry.Info("request started")
		// Output: time="2019-11-11T20:47:29.5326335+01:00" level=info msg="request started" app=example reqID=5577006791947779410

		// setting contextual log entry for the handler
		ctx := logctx.New(r.Context(), logEntry)
		next.ServeHTTP(w, r.WithContext(ctx))

		logEntry.Info("request finished")
		// Output: time="2019-11-11T20:47:29.5746316+01:00" level=info msg="request finished" app=example reqID=5577006791947779410
	})
}
