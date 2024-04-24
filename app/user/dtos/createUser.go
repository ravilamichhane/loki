package dtos

import (
	"app/user/entities"
	"nest/thor/validate"
)

type CreateUser struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateUser) Validate() error {
	return validate.Check(c)
}

func (c CreateUser) ToUser() *entities.User {
	return &entities.User{
		Name: c.Name,
	}
}
