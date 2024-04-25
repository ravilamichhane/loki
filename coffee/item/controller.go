package item

import (
	"coffee/item/dtos"
	"nest/common"
	"nest/thor"
)

type ItemController struct {
	ItemService ItemService
}

func NewItemController(ItemService ItemService) *ItemController {
	return &ItemController{
		ItemService: ItemService,
	}
}

func (u *ItemController) Routes() []common.Route {
	return []common.Route{
		common.GET("/", u.FindAll),
		common.GET("/{id}", u.FindOne),
		common.POST("", u.Create),
		common.PUT("/{id}", u.Update),
		common.DELETE("/{id}", u.Delete),
	}
}

func (u *ItemController) FindAll(ctx common.HttpContext) error {
	ctx.SetStatusCode(200)

	items, err := u.ItemService.FindAll()

	if err != nil {
		return err
	}

	ctx.JSON(200, items)
	return nil
}

func (u *ItemController) FindOne(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	item, err := u.ItemService.FindOne(uid)

	if err != nil {
		return err
	}

	ctx.SetStatusCode(200)

	ctx.JSON(200, item)
	return nil
}

func (u *ItemController) Create(ctx common.HttpContext) error {
	var item dtos.CreateItem

	if err := ctx.Decode(&item); err != nil {

		return err
	}

	newItem := item.ToItem()

	if err := u.ItemService.Create(newItem); err != nil {
		return err
	}

	ctx.JSON(200, newItem)
	return nil

}

func (u *ItemController) Update(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	item, err := u.ItemService.FindOne(uid)

	if err != nil {
		return err
	}

	var updateItem dtos.UpdateItem

	if err := ctx.Decode(&updateItem); err != nil {
		return err
	}
	updateItem.Decode(item)

	if err = u.ItemService.Update(item); err != nil {
		return err
	}

	ctx.JSON(200, item)
	return nil
}

func (u *ItemController) Delete(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}
	err = u.ItemService.Delete(uid)

	if err == nil {
		ctx.JSON(200, struct {
			Message string `json:"message"`
		}{
			Message: "Item Deleted",
		})
	}

	return err

}

func (u *ItemController) Prefix() string {
	return "/api/v1/item"
}

func (u *ItemController) Middlewares() []common.MiddleWare {
	return []common.MiddleWare{
		thor.LoggingMiddleware,
		thor.ErrorMiddleware,
		thor.PanicMiddleWare,
	}
}
