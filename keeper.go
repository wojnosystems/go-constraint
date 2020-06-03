package constraint

import (
	"github.com/wojnosystems/go-path/go_path"
)

// Keepers keep a set of constrains from a request to validate a structure
type ViolationContainer interface {
	// IsEmpty is true if no infringements were added
	IsEmpty() bool

	// GetPaths gets a list of Paths that were appended, each path has one or more infringements
	GetPaths() []go_path.Pather

	// Get the list of infringements at a Path
	Get(go_path.Pather) []Infringer

	// Len is the number of errors across all paths
	Len() int
}

// ViolationAppender allows infringements to be added
type ViolationAppender interface {
	// Append an infringement at the specified path, this is how you add infringements to the set
	Append(go_path.Pather, Infringer)
}

type ViolationMerger interface {
	// Merge the violations in one ViolationContainer into the caller.
	// The violations' current path is appended to the path provided by the caller
	// so that the violations appear to be children of that path
	Merge(go_path.Pather, ViolationContainer) ViolationMutater
}

type ViolationMutater interface {
	ViolationContainer
	ViolationAppender
	ViolationMerger
}
