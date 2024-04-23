package dtos

import (
	"app/user/entities"
	"fmt"
	"nest/thor"
)

type CreateUser struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateUser) Validate() error {
	if c.Name == "" {
		return thor.NewFieldError("name", fmt.Errorf("name is required"))
	}
	return nil
}

func (c CreateUser) ToUser() *entities.User {
	return &entities.User{
		Name: c.Name,
	}
}
