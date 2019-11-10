package errfield_test

import (
	"errors"
	"os"

	"github.com/go-logrusutil/logrusutil/errfield"

	log "github.com/sirupsen/logrus"
)

func Example_basic() {
	log.SetOutput(os.Stdout)

	// setup the errfield.Formatter
	log.SetFormatter(&errfield.Formatter{
		Formatter: &log.TextFormatter{DisableTimestamp: true},
	})

	// use errfield.Add to add fields
	err := errors.New("something failed")
	err = errfield.Add(err, "foo", "bar")
	log.WithError(err).Error("crash")
	// Output: level=error msg=crash error="something failed" foo=bar
}

func ExampleFormatter_errorFieldsKey() {
	log.SetOutput(os.Stdout)

	// setup the errfield.Formatter with ErrorFieldsKey
	log.SetFormatter(&errfield.Formatter{
		Formatter:      &log.TextFormatter{DisableTimestamp: true},
		ErrorFieldsKey: "error_fields",
	})

	// use errfield.Add to add fields
	err := errors.New("something failed")
	err = errfield.Add(err, "fizz", "buzz")
	log.WithError(err).Error("crash")
	// Output: level=error msg=crash error="something failed" error_fields="map[fizz:buzz]"
}
