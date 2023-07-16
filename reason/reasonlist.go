package reason

import (
	"errors"
	"sort"
)

// ReasonList represents an ordered list of reasons.
type ReasonList []Reason

// Ensure [ReasonList] implements the [error] interface.
var _ error = ReasonList{}

// Add allows a reason to be added to the list.
func (rl *ReasonList) Add(reason Reason) {
	*rl = append(*rl, reason)
}

// Error reports an error string of all underlying errors.
func (rl ReasonList) Error() string {
	return errors.Join(rl.Unwrap()...).Error()
}

// Unwrap returns a list of all the wrapped errors.
func (rl ReasonList) Unwrap() []error {
	sort.Slice(rl, func(i, j int) bool {
		return rl[i].priority > rl[j].priority
	})

	errs := make([]error, 0, len(rl))

	for _, reason := range rl {
		errs = append(errs, reason.Unwrap())
	}

	return errs
}
