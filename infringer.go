package constraint

// Infinger holds the details of how a constraint was violated or infringed
type Infringer interface {
	// Name is the human-readable version of how the constraint was violated, or infringed
	// e.g. This could be "length too long" or "end time before start time"
	Name() string

	// Remedy is an explanation about how the constraint can be satisfied, or,
	// how the violation or infringement can be resolved or remedied
	// e.g. "remove 5 characters" or "start time needs to be before the end time"
	Remedy() string
}
