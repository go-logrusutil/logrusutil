package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-logrusutil/logrusutil/errfield"
	"github.com/go-logrusutil/logrusutil/logctx"
)

func main() {
	// setup logging
	log.SetOutput(os.Stdout)
	// set errfield.Formatter to log the error fields
	log.SetFormatter(&errfield.Formatter{
		Formatter: &log.TextFormatter{
			TimestampFormat: time.RFC3339Nano,
		},
	})
	// set logctx.Default and use it in places when there is no context
	logctx.Default = log.WithField("app", "example")
	logctx.Default.Info("server starting")
	// Output: time="2019-11-11T20:45:01.3777602+01:00" level=info msg="server starting" app=example

	// bootstrap HTTP server
	mux := http.NewServeMux()
	registerHello(mux)
	registerTry(mux)
	registerGo(mux)
	// create a server with structred contextual logging middleware
	server := http.Server{
		Addr:    ":8080",
		Handler: logMiddleware(mux),
	}
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		logctx.Default.WithError(err).Fatal("server unexpectedly closed")
	}
	logctx.Default.Info("server closed")
}
