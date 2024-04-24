package dtos

import (
	"app/user/entities"
	"nest/thor/validate"
)

type UpdateUser struct {
	Name string `json:"name"`
}

func (u UpdateUser) Validate() error {
	return validate.Check(u)
}

func (u UpdateUser) ToUser(user *entities.User) *entities.User {
	if u.Name != "" {
		user.Name = u.Name
	}
	return user
}
