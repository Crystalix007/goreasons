package reasons_test

import (
	"errors"
	"testing"

	"github.com/Crystalix007/goreasons/reason"
	"github.com/Crystalix007/goreasons/reasons"
	"github.com/stretchr/testify/assert"
)

func TestReasons(t *testing.T) {
	t.Parallel()

	reasons := reasons.Reasons{}

	nonFatalErr := errors.New("reasons_test: non-fatal error")
	reasons.NonFatal("test message", nonFatalErr)

	assert.ErrorIs(t, reasons, nonFatalErr)

	fatalErr := errors.New("reasons_test: fatal error")

	assert.ErrorIs(t, reasons.Fatal("test message", fatalErr), fatalErr)

	// Should report all errors.
	assert.ErrorIs(t, reasons, fatalErr)
	assert.ErrorIs(t, reasons, nonFatalErr)
}

func TestReasonsReport(t *testing.T) {
	t.Parallel()

	rs := reasons.Reasons{}
	err := errors.New("reasons_test: test error")

	assert.NoError(t, rs.Report(reasons.NON_FATAL, "non-fatal message", err, reason.Priority(reason.HIGH)))
	assert.ErrorIs(t, rs.Report(reasons.FATAL, "fatal message", err, reason.Priority(reason.LOW)), err)

	assert.ErrorIs(t, rs.Report(reasons.Severity(20), "unknown severity", err), reasons.ErrUnknownSeverity)
}

func TestReasonsFailed(t *testing.T) {
	t.Parallel()

	rs := reasons.Reasons{}

	// Shouldn't report to have failed without an issue.
	assert.False(t, rs.Failed())

	err := errors.New("reasons_test: test non error")

	// Should report to have failed if there is any non-fatal error.
	rs.NonFatal("non-fatal error", err)
	assert.True(t, rs.Failed())

	rs = reasons.Reasons{}

	// Should report to have failed if there is any fatal error.
	rs.Fatal("fatal error", err)
	assert.True(t, rs.Failed())

	// Should report to have failed if there are both non-fatal and fatal
	// errors.
	rs.NonFatal("non-fatal error", err)
	assert.True(t, rs.Failed())
}

func TestReasonsErrorContainsErrors(t *testing.T) {
	t.Parallel()

	reasons := reasons.Reasons{}

	nonFatalErr := errors.New("reasons_test: non-fatal error")
	reasons.NonFatal("non-fatal test message", nonFatalErr)

	fatalErr := errors.New("reasons_test: fatal error")
	reasons.Fatal("fatal test message", fatalErr)

	assert.ErrorContains(t, reasons,
		`Fatal errors:

fatal test message: reasons_test: fatal error

Non-fatal errors:

non-fatal test message: reasons_test: non-fatal error`)
}
