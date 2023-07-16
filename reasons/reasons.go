// Package reasons provides a way to aggregate non-fatal and fatal errors, to
// provide more detailed traces.
package reasons

import (
	"errors"
	"strings"

	"github.com/Crystalix007/goreasons/reason"
)

// UnknownErrorSeverity is returned when an unknown error severity is provided.
var UnknownErrorSeverity = errors.New("reasons: unknown error severity")

// Reasons aggregates non-fatal errors and fatal errors.
type Reasons struct {
	// nonFatal reasons are collected, but will not cause a failure.
	nonFatal reason.ReasonList

	// fatal reasons are collected, and return an immediate error.
	fatal reason.ReasonList
}

// Ensure [Reasons] fulfills the [error] interface.
var _ error = Reasons{}

// Report reports an error with the provided severity and reason. Optional
// options allow further configuration of the reason.
func (r *Reasons) Report(
	severity Severity,
	message string,
	err error,
	opts ...reason.Option,
) error {
	switch severity {
	case NON_FATAL:
		r.NonFatal(message, err, opts...)

		return nil
	case FATAL:
		return r.Fatal(message, err, opts...)
	}

	return UnknownErrorSeverity
}

// NonFatal allows a non-fatal error to be reported.
func (r *Reasons) NonFatal(
	message string,
	err error,
	opts ...reason.Option,
) {
	r.nonFatal = append(r.nonFatal, reason.New(message, err, opts...))
}

// Fatal allows a fatal error to be reported. Returns the current reasons
// as an error.
func (r *Reasons) Fatal(
	message string,
	err error,
	opts ...reason.Option,
) error {
	r.fatal = append(r.fatal, reason.New(message, err, opts...))

	return r
}

// Error returns the error string of all underlying errors.
// Returns the empty string if there are no underlying errors.
// Returns first fatal errors (if any), then non-fatal errors (if any).
func (r Reasons) Error() string {
	var errString strings.Builder

	if len(r.fatal) > 0 {
		errString.WriteString("Fatal errors:\n\n")
		errString.WriteString(r.fatal.Error())

		if len(r.nonFatal) > 0 {
			errString.WriteString("\n\n")
		}
	}

	if len(r.fatal) > 0 {
		errString.WriteString("Non-fatal errors:\n\n")
		errString.WriteString(r.nonFatal.Error())
	}

	return errString.String()
}

// Unwrap returns all the wrapped errors.
func (r Reasons) Unwrap() []error {
	return append(r.fatal.Unwrap(), r.nonFatal.Unwrap()...)
}
