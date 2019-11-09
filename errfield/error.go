package errfield

import "errors"

// Error contains custom fields and wrapped error.
type Error struct {
	Err    error
	Fields map[string]interface{}
}

// Unwrap returns the wrapped error.
func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Error() string {
	return e.Err.Error()
}

// Add return error and adds a field to *errfield.Error in the error chain.
func Add(err error, key string, value interface{}) error {
	var e *Error
	if errors.As(err, &e) {
		e.Fields[key] = value
		return err
	}
	return &Error{
		Err:    err,
		Fields: map[string]interface{}{key: value},
	}
}
