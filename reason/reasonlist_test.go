package reason_test

import (
	"errors"
	"testing"

	"github.com/Crystalix007/goreasons/reason"
	"github.com/stretchr/testify/assert"
)

func TestReasonList(t *testing.T) {
	t.Parallel()

	rl := reason.ReasonList{}
	testErr := errors.New("reason_test: test error")

	rl.Add(reason.New("test message", testErr))

	assert.ErrorIs(t, rl, testErr)
}

func TestReasonListSorting(t *testing.T) {
	t.Parallel()

	rl := reason.ReasonList{}
	testErr1 := errors.New("reason_test: test error 1")
	testErr2 := errors.New("reason_test: test error 2")

	rl.Add(reason.New("test message 1", testErr1, reason.Priority(reason.LOW)))
	rl.Add(reason.New("test message 2", testErr2, reason.Priority(reason.HIGH)))

	errs := rl.Unwrap()

	// Sorts priority high to low.
	assert.ErrorIs(t, errs[0], testErr2)
	assert.ErrorIs(t, errs[1], testErr1)
}

func TestReasonListError(t *testing.T) {
	t.Parallel()

	rl := reason.ReasonList{}
	testErr1 := errors.New("reason_test: test error 1")
	testErr2 := errors.New("reason_test: test error 2")

	rl.Add(reason.New("test message 1", testErr1))
	rl.Add(reason.New("test message 2", testErr2))

	assert.ErrorContains(t, rl, testErr1.Error())
	assert.ErrorContains(t, rl, testErr2.Error())
}
