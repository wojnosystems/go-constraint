package constraint_provider

import (
	"bytes"
	"github.com/wojnosystems/go-constraint"
	"github.com/wojnosystems/go-path/go_path"
)

type violations struct {
	violationSum int
	collection   map[string][]constraint.Infringer
}

func NewViolations() constraint.KeepMutator {
	return &violations{
		violationSum: 0,
		collection:   make(map[string][]constraint.Infringer),
	}
}

func (v violations) IsEmpty() bool {
	return len(v.collection) == 0
}

func (v violations) Len() int {
	return v.violationSum
}

func (v violations) GetPaths() (paths []go_path.Pather) {
	paths = make([]go_path.Pather, 0, len(v.collection))
	for key := range v.collection {
		path, _ := go_path.Parse(bytes.NewBufferString(key))
		paths = append(paths, path)
	}
	return
}

func (v violations) Get(path go_path.Pather) (violations []constraint.Infringer) {
	key := path.String()
	if violationList, ok := v.collection[key]; ok {
		violations = make([]constraint.Infringer, 0, len(violationList))
		for _, violation := range violationList {
			violations = append(violations, violation)
		}
	} else {
		violations = make([]constraint.Infringer, 0)
	}
	return
}

func (v *violations) Append(path go_path.Pather, violation constraint.Infringer) {
	key := path.String()
	if _, ok := v.collection[key]; !ok {
		v.collection[key] = make([]constraint.Infringer, 0, 1)
	}
	v.collection[key] = append(v.collection[key], violation)
	v.violationSum++
}

func (v *violations) Merge(tailingPath go_path.Pather, mergeFrom constraint.ViolationContainer) (new constraint.KeepMutator) {
	new = NewViolations()
	for key, violaters := range v.collection {
		path, _ := go_path.Parse(bytes.NewBufferString(key))
		for _, violater := range violaters {
			new.Append(path, violater)
		}
	}
	for _, path := range mergeFrom.GetPaths() {
		infringements := mergeFrom.Get(path)
		insertPath := tailingPath.Copy()
		path.Each(func(index int, componenter go_path.Componenter) {
			insertPath.Append(componenter)
		})
		for _, infringement := range infringements {
			new.Append(insertPath, infringement)
		}
	}
	return
}
