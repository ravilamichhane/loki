package dtos

import (
	"auth/user/entities"
	"loki/thor/validate"
)

type CreateUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (c CreateUser) Validate() error {
	return validate.Check(c)
}

func (c CreateUser) ToUser() *entities.User {
	return &entities.User{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Password:  c.Password,
	}
}
