package main

import (
	"net/http"

	"github.com/go-logrusutil/logrusutil/logctx"
)

func registerHello(mux *http.ServeMux) {
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// log with contextual data
		logctx.From(r.Context()).WithField("foo", "bar").Info("hello world")
		// Output: time="2019-11-11T20:47:29.5366318+01:00" level=info msg="hello world" app=example foo=bar reqID=5577006791947779410
	})
}
