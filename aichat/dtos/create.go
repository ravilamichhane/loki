package dtos

import (
	"aichat/entities"
	"nest/thor/validate"
)

type CreateAichat struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateAichat) Validate() error {
	return validate.Check(c)
}

func (c CreateAichat) ToAichat() *entities.Aichat {
	return &entities.Aichat{
		Name: c.Name,
	}
}
