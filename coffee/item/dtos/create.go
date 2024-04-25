package dtos

import (
	"coffee/item/entities"
	"nest/thor/validate"
)

type CreateItem struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateItem) Validate() error {
	return validate.Check(c)
}

func (c CreateItem) ToItem() *entities.Item {
	return &entities.Item{
		Name: c.Name,
	}
}
