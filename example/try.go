package main

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/go-logrusutil/logrusutil/errfield"
	"github.com/go-logrusutil/logrusutil/logctx"
)

func registerTry(mux *http.ServeMux) {
	mux.HandleFunc("/try", func(w http.ResponseWriter, r *http.Request) {
		if err := try(); err != nil {
			// log error with contextual data
			logctx.From(r.Context()).WithError(err).Error("try failed")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logctx.From(r.Context()).Info("try succeded")
	})
}

func try() error {
	p := struct {
		X int
		Y int
	}{
		rand.Intn(4),
		rand.Intn(4),
	}
	if p.X != p.Y {
		// return error with a structured field
		return errfield.Add(errors.New("failed to generate an excelent point"), "point", p)
	}
	return nil
}
