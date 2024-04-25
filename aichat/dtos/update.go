package dtos

import (
	"aichat/entities"
	"nest/thor/validate"
)

type UpdateAichat struct {
	Name string `json:"name"`
}

func (u UpdateAichat) Validate() error {
	return validate.Check(u)
}

func (u UpdateAichat) Decode(aichat *entities.Aichat)  {
	if u.Name != "" {
		aichat.Name = u.Name
	}
}
