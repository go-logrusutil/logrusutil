package main

import (
	"math/rand"
	"net/http"

	"github.com/go-logrusutil/logrusutil/logctx"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := rand.Int()
		logEntry := logctx.DefaultLogEntry.WithField("request_id", reqID)
		logEntry.Info("request started")

		// setting contextual log entry for the handler
		ctx := logctx.New(r.Context(), logEntry)
		next.ServeHTTP(w, r.WithContext(ctx))

		logEntry.Info("request finished")
	})
}
