package dtos

import (
	"auth/todo/entities"
	"loki/thor/validate"
)

type CreateTodo struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateTodo) Validate() error {
	return validate.Check(c)
}

func (c CreateTodo) ToTodo() *entities.Todo {
	return &entities.Todo{
		Name: c.Name,
	}
}
