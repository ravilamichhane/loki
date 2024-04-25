package shop

import (
	"coffee/shop/dtos"
	"nest/common"
	"nest/thor"
)

type ShopController struct {
	ShopService ShopService
}

func NewShopController(ShopService ShopService) *ShopController {
	return &ShopController{
		ShopService: ShopService,
	}
}

func (u *ShopController) Routes() []common.Route {
	return []common.Route{
		common.GET("/", u.FindAll),
		common.GET("/{id}", u.FindOne),
		common.POST("", u.Create),
		common.PUT("/{id}", u.Update),
		common.DELETE("/{id}", u.Delete),
	}
}

func (u *ShopController) FindAll(ctx common.HttpContext) error {
	ctx.SetStatusCode(200)

	shops, err := u.ShopService.FindAll()

	if err != nil {
		return err
	}

	ctx.JSON(200, shops)
	return nil
}

func (u *ShopController) FindOne(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	shop, err := u.ShopService.FindOne(uid)

	if err != nil {
		return err
	}

	ctx.SetStatusCode(200)

	ctx.JSON(200, shop)
	return nil
}

func (u *ShopController) Create(ctx common.HttpContext) error {
	var shop dtos.CreateShop

	if err := ctx.Decode(&shop); err != nil {

		return err
	}

	newShop := shop.ToShop()

	if err := u.ShopService.Create(newShop); err != nil {
		return err
	}

	ctx.JSON(200, newShop)
	return nil

}

func (u *ShopController) Update(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	shop, err := u.ShopService.FindOne(uid)

	if err != nil {
		return err
	}

	var updateShop dtos.UpdateShop

	if err := ctx.Decode(&updateShop); err != nil {
		return err
	}
	updateShop.Decode(shop)

	if err = u.ShopService.Update(shop); err != nil {
		return err
	}

	ctx.JSON(200, shop)
	return nil
}

func (u *ShopController) Delete(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}
	err = u.ShopService.Delete(uid)

	if err == nil {
		ctx.JSON(200, struct {
			Message string `json:"message"`
		}{
			Message: "Shop Deleted",
		})
	}

	return err

}

func (u *ShopController) Prefix() string {
	return "/api/v1/shop"
}

func (u *ShopController) Middlewares() []common.MiddleWare {
	return []common.MiddleWare{
		thor.LoggingMiddleware,
		thor.ErrorMiddleware,
		thor.PanicMiddleWare,
	}
}
