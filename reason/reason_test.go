package reason_test

import (
	"errors"
	"testing"

	"github.com/Crystalix007/goreasons/reason"
	"github.com/stretchr/testify/assert"
)

func TestReasonIncludesMessage(t *testing.T) {
	t.Parallel()

	testMessage := "test message"

	assert.ErrorContains(t, reason.New(testMessage, nil), testMessage)
}

func TestReasonIsWrappedError(t *testing.T) {
	t.Parallel()

	testReason := errors.New("reasons_test: test wrapped error")

	assert.ErrorIs(t, reason.New("test message", testReason), testReason)
}
