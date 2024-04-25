package aichat

import (
	"aichat/dtos"
	"nest/common"
	"nest/thor"
)

type AichatController struct {
	AichatService AichatService
}

func NewAichatController(AichatService AichatService) *AichatController {
	return &AichatController{
		AichatService: AichatService,
	}
}

func (u *AichatController) Routes() []common.Route {
	return []common.Route{
		common.GET("/", u.FindAll),
		common.GET("/{id}", u.FindOne),
		common.POST("", u.Create),
		common.PUT("/{id}", u.Update),
		common.DELETE("/{id}", u.Delete),
	}
}

func (u *AichatController) FindAll(ctx common.HttpContext) error {
	ctx.SetStatusCode(200)

	aichats, err := u.AichatService.FindAll()

	if err != nil {
		return err
	}

	ctx.JSON(200, aichats)
	return nil
}

func (u *AichatController) FindOne(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	aichat, err := u.AichatService.FindOne(uid)

	if err != nil {
		return err
	}

	ctx.SetStatusCode(200)

	ctx.JSON(200, aichat)
	return nil
}

func (u *AichatController) Create(ctx common.HttpContext) error {
	var aichat dtos.CreateAichat

	if err := ctx.Decode(&aichat); err != nil {

		return err
	}

	newAichat := aichat.ToAichat()

	if err := u.AichatService.Create(newAichat); err != nil {
		return err
	}

	ctx.JSON(200, newAichat)
	return nil

}

func (u *AichatController) Update(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	aichat, err := u.AichatService.FindOne(uid)

	if err != nil {
		return err
	}

	var updateAichat dtos.UpdateAichat

	if err := ctx.Decode(&updateAichat); err != nil {
		return err
	}
	updateAichat.Decode(aichat)

	if err = u.AichatService.Update(aichat); err != nil {
		return err
	}

	ctx.JSON(200, aichat)
	return nil
}

func (u *AichatController) Delete(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}
	err = u.AichatService.Delete(uid)

	if err == nil {
		ctx.JSON(200, struct {
			Message string `json:"message"`
		}{
			Message: "Aichat Deleted",
		})
	}

	return err

}

func (u *AichatController) Prefix() string {
	return "/api/v1/aichat"
}

func (u *AichatController) Middlewares() []common.MiddleWare {
	return []common.MiddleWare{
		thor.LoggingMiddleware,
		thor.ErrorMiddleware,
		thor.PanicMiddleWare,
	}
}
