package constraint_provider

import "github.com/wojnosystems/go-constraint"

type violation struct {
	name   string
	remedy string
}

func NewViolation(name string, remedy string) constraint.Infringer {
	return &violation{
		name:   name,
		remedy: remedy,
	}
}

func (v *violation) Name() string {
	return v.name
}
func (v *violation) Remedy() string {
	return v.remedy
}
