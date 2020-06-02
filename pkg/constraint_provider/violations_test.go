package constraint_provider

import (
	"github.com/stretchr/testify/assert"
	"github.com/wojnosystems/go-constraint"
	"github.com/wojnosystems/go-path/go_path"
	"testing"
)

func TestViolations_Append(t *testing.T) {
	cases := map[string]struct {
		setUp             func() constraint.KeepMutator
		expectedPathCount int
		expectedLen       int
	}{
		"do nothing": {
			setUp: func() constraint.KeepMutator {
				return NewViolations()
			},
		},
		"one": {
			setUp: func() constraint.KeepMutator {
				v := NewViolations()
				r := go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test"))
				v.Append(r, NewViolation("test"))
				return v
			},
			expectedPathCount: 1,
			expectedLen:       1,
		},
		"two": {
			setUp: func() constraint.KeepMutator {
				v := NewViolations()
				r := go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test"))
				v.Append(r, NewViolation("test"))
				r = go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test2"))
				v.Append(r, NewViolation("test2"))
				return v
			},
			expectedPathCount: 2,
			expectedLen:       2,
		},
		"two same path": {
			setUp: func() constraint.KeepMutator {
				v := NewViolations()
				r := go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test"))
				v.Append(r, NewViolation("test"))
				v.Append(r, NewViolation("test2"))
				return v
			},
			expectedPathCount: 1,
			expectedLen:       2,
		},
	}

	for caseName, c := range cases {
		actual := c.setUp()
		paths := actual.GetPaths()
		assert.Equal(t, c.expectedPathCount, len(paths), caseName)
		assert.Equal(t, c.expectedLen, actual.Len(), caseName)
	}
}
