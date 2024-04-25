package dtos

import (
	"coffee/shop/entities"
	"nest/thor/validate"
)

type CreateShop struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateShop) Validate() error {
	return validate.Check(c)
}

func (c CreateShop) ToShop() *entities.Shop {
	return &entities.Shop{
		Name: c.Name,
	}
}
