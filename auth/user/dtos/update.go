package dtos

import (
	"auth/user/entities"
	"loki/thor/validate"
)

type UpdateUser struct {
	FistName string `json:"first_name" `
	LastName string `json:"last_name" `
	Email    string `json:"email" `
}

func (u UpdateUser) Validate() error {
	return validate.Check(u)
}

func (u UpdateUser) Decode(user *entities.User) {
	if u.FistName != "" {
		user.FirstName = u.FistName
	}
	if u.LastName != "" {
		user.LastName = u.LastName
	}

	if u.Email != "" {
		user.Email = u.Email
	}
}
