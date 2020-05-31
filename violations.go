package constraint

type Violations interface {
	IsEmpty() bool
	GetPaths() []Pather
	Get(Pather) []Violation
}

type ViolationsAppender interface {
	Append(Pather, Violation)
}

type ViolationsMerger interface {
	Merge(Pather, Violations) Violations
}

type ViolationsMutator interface {
	ViolationsAppender
}
