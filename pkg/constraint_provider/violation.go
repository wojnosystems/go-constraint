package constraint_provider

import "github.com/wojnosystems/go-constraint"

type violation struct {
	summary string
	remedy  string
}

func NewViolation(summary string, remedy string) constraint.Infringer {
	return &violation{
		summary: summary,
		remedy:  remedy,
	}
}

func (v *violation) Summary() string {
	return v.summary
}
func (v *violation) Remedy() string {
	return v.remedy
}
