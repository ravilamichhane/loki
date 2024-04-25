package dto

import (
	"auth/user/dtos"
	"nest/thor/validate"
)

type SignUpRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (s *SignUpRequest) Validate() error {
	return validate.Check(s)
}

func (s *SignUpRequest) ToCreateUser() dtos.CreateUser {
	return dtos.CreateUser{
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		Password:  s.Password,
	}
}
