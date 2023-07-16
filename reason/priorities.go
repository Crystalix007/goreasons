package reason

// ReasonPriority represents one taxonomy of priorities, although any suitable
// scale may be substituted.
type ReasonPriority = uint

const (
	// LOW represents a low-priority reason that is unlikely to
	// materially change whether the operation fails.
	LOW = 1

	// MEDIUM represents a medium-priority reason that may or may
	// materially change whether the operation fails.
	MEDIUM = 10

	// HIGH represents a high-priority reason that is likely to
	// materially change whether the operation fails.
	HIGH = 100
)
