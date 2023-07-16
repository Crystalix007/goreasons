// Package reason implements a wrapping error type, to allow futher properties
// to be stored on an error.
package reason

import "fmt"

// Reason represents a single error and associated metadata.
type Reason struct {
	message  string
	err      error
	priority ReasonPriority
}

// Ensure [Reason] implements the [error] type.
var _ error = Reason{}

// New constructs a new reason from the provided message and error.
// Optional options can also be provided to further customise the reason.
func New(message string, err error, opts ...Option) Reason {
	reason := Reason{
		message: message,
		err:     err,
	}

	for _, opt := range opts {
		opt(&reason)
	}

	return reason
}

// Error returns the error string of the underlying error.
func (r Reason) Error() string {
	return r.Unwrap().Error()
}

// Unwrap returns the underlying error.
func (r Reason) Unwrap() error {
	return fmt.Errorf("%s: %w", r.message, r.err)
}
