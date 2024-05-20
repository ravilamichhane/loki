package dtos

import (
	"auth/todo/entities"
	"nest/thor/validate"
)

type UpdateTodo struct {
	Name string `json:"name"`
}

func (u UpdateTodo) Validate() error {
	return validate.Check(u)
}

func (u UpdateTodo) Decode(todo *entities.Todo)  {
	if u.Name != "" {
		todo.Name = u.Name
	}
}
