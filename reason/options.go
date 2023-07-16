package reason

// Option represents an option that can be set on a [Reason].
type Option func(reason *Reason)

// Priority sets a [ReasonPriority] on a [Reason].
func Priority(priority ReasonPriority) Option {
	return func(reason *Reason) {
		reason.priority = priority
	}
}
