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
	// set logctx.DefaultLogEntry and use it in places when there is no context
	logctx.DefaultLogEntry = log.WithField("app", "example")
	logctx.DefaultLogEntry.Info("server starting")

	// bootstrap HTTP server
	mux := http.NewServeMux()
	registerHello(mux)
	registerTry(mux)
	// create a server with structred contextual logging middleware
	server := http.Server{
		Addr:    ":8080",
		Handler: logMiddleware(mux),
	}
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		logctx.DefaultLogEntry.WithError(err).Fatal("server unexpectedly closed")
	}
	logctx.DefaultLogEntry.Info("server closed")
}
