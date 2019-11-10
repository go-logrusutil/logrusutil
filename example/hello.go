package main

import (
	"net/http"

	"github.com/go-logrusutil/logrusutil/logctx"
)

func registerHello(mux *http.ServeMux) {
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// log with contextual data
		logctx.From(r.Context()).WithField("foo", "bar").Info("hello world")
	})
}
