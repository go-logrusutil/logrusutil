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
			// Output: time="2019-11-11T20:49:50.7667768+01:00" level=error msg="try failed" app=example error="failed to generate an excelent point" point="{2 1}" reqID=3916589616287113937
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logctx.From(r.Context()).Info("try succeded")
		// Output: time="2019-11-11T20:48:27.2073124+01:00" level=info msg="try succeded" app=example reqID=8674665223082153551
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
