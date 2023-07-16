package reasons

// Severity provides a classification of error severities.
type Severity uint

const (
	// NON_FATAL represents an error that shouldn't cause immediate
	// failure.
	NON_FATAL Severity = iota

	// FATAL represents an error that should cause immediate failure.
	FATAL
)
