package constraint_provider

import "github.com/wojnosystems/constraint"

type violations struct {
	collection map[constraint.Pather][]constraint.Violation
}

func NewViolations() constraint.Violations {
	return &violations{
		collection: make(map[constraint.Pather][]constraint.Violation),
	}
}

func (v violations) IsEmpty() bool {
	return len(v.collection) == 0
}

func (v violations) GetPaths() (paths []constraint.Pather) {
	paths = make([]constraint.Pather, 0, len(v.collection))
	for key := range v.collection {
		paths = append(paths, key)
	}
	return
}

func (v violations) Get(path constraint.Pather) (violations []constraint.Violation) {
	if violationList, ok := v.collection[path]; ok {
		violations = make([]constraint.Violation, 0, len(violationList))
		for _, violation := range violationList {
			violations = append(violations, violation)
		}
	} else {
		violations = make([]constraint.Violation, 0)
	}
	return
}

func (v *violations) Append(path constraint.Pather, violation constraint.Violation) {
	if _, ok := v.collection[path]; !ok {
		v.collection[path] = make([]constraint.Violation, 0, 1)
	}
	v.collection[path] = append(v.collection[path], violation)
}
