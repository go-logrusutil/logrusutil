// Package errfield adds possibility to wrap errors with fields and then log them in structured way.
package errfield

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// Formatter decorates logrus.Formatter to add error fields under to the log entry.
// Implements logrus.Formatter.
type Formatter struct {
	// Formatter is the decorated logrus.Formatter.
	// Default TextFormatter is used when none provided.
	logrus.Formatter

	// ErrorFieldsKey defines under which key the error log fields would be added.
	// For empty string it des not create any dedicated key.
	ErrorFieldsKey string
}

// Format implements logrus.Formatter.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var e *Error
	err, ok := entry.Data[logrus.ErrorKey].(error)
	if ok && errors.As(err, &e) {
		destErrorFields := f.errorFields(entry)
		for key, value := range e.Fields {
			destErrorFields[key] = value
		}
	}
	return f.baseFormatter().Format(entry)
}

func (f *Formatter) baseFormatter() logrus.Formatter {
	if f.Formatter == nil {
		return &logrus.TextFormatter{}
	}
	return f.Formatter
}

func (f *Formatter) errorFields(entry *logrus.Entry) map[string]interface{} {
	if f.ErrorFieldsKey == "" {
		return entry.Data
	}
	res := map[string]interface{}{}
	entry.Data[f.ErrorFieldsKey] = res
	return res
}
