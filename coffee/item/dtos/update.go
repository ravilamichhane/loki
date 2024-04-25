package dtos

import (
	"coffee/item/entities"
	"nest/thor/validate"
)

type UpdateItem struct {
	Name string `json:"name"`
}

func (u UpdateItem) Validate() error {
	return validate.Check(u)
}

func (u UpdateItem) Decode(item *entities.Item)  {
	if u.Name != "" {
		item.Name = u.Name
	}
}
