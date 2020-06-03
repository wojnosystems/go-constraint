package constraint_provider

import (
	"github.com/stretchr/testify/assert"
	"github.com/wojnosystems/go-constraint"
	"github.com/wojnosystems/go-path/go_path"
	"testing"
)

func TestViolations_Append(t *testing.T) {
	name := "name"
	remedy := "remedy"
	cases := map[string]struct {
		setUp             func() constraint.ViolationMutater
		expectedPathCount int
		expectedLen       int
	}{
		"do nothing": {
			setUp: func() constraint.ViolationMutater {
				return NewViolations()
			},
		},
		"one": {
			setUp: func() constraint.ViolationMutater {
				v := NewViolations()
				r := go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test"))
				v.Append(r, NewViolation(name, remedy))
				return v
			},
			expectedPathCount: 1,
			expectedLen:       1,
		},
		"two": {
			setUp: func() constraint.ViolationMutater {
				v := NewViolations()
				r := go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test"))
				v.Append(r, NewViolation(name, remedy))
				r = go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test2"))
				v.Append(r, NewViolation(name, remedy))
				return v
			},
			expectedPathCount: 2,
			expectedLen:       2,
		},
		"two same path": {
			setUp: func() constraint.ViolationMutater {
				v := NewViolations()
				r := go_path.NewRoot()
				r.Append(go_path.NewInstanceVariableNamed("test"))
				v.Append(r, NewViolation(name, remedy))
				v.Append(r, NewViolation(name, remedy))
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
