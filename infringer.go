package constraint

// Infinger holds the details of how a constraint was violated or infringed
type Infringer interface {
	// Summary is the human-readable version of how the constraint was violated, or infringed
	Summary() string

	// Remedy is how the constraint can be satisfied, or, how the violation or infringement can be resolved or remedied
	Remedy() string
}
