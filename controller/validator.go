package controller

import (
	"gopkg.in/go-playground/validator.v9"
)

var v = validator.New()

type ProblemValidator struct {
	validator *validator.Validate
}

func (p *ProblemValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}
