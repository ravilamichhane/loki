package dtos

import (
	"coffee/shop/entities"
	"nest/thor/validate"
)

type UpdateShop struct {
	Name string `json:"name"`
}

func (u UpdateShop) Validate() error {
	return validate.Check(u)
}

func (u UpdateShop) Decode(shop *entities.Shop)  {
	if u.Name != "" {
		shop.Name = u.Name
	}
}
