package user

import (
	"auth/user/dtos"
	"nest/common"
	"nest/thor"
)

type UserController struct {
	UserService UserService
}

func NewUserController(UserService UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (u *UserController) Routes() []common.Route {
	return []common.Route{
		common.GET("/", u.FindAll),
		common.GET("/{id}", u.FindOne),
		common.POST("", u.Create),
		common.PUT("/{id}", u.Update),
		common.DELETE("/{id}", u.Delete),
	}
}

func (u *UserController) FindAll(ctx common.HttpContext) error {
	ctx.SetStatusCode(200)

	users, err := u.UserService.FindAll()

	if err != nil {
		return err
	}

	ctx.JSON(200, users)
	return nil
}

func (u *UserController) FindOne(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	user, err := u.UserService.FindOne(uid)

	if err != nil {
		return err
	}

	ctx.SetStatusCode(200)

	ctx.JSON(200, user)
	return nil
}

func (u *UserController) Create(ctx common.HttpContext) error {
	var user dtos.CreateUser

	if err := ctx.Decode(&user); err != nil {
		return err
	}

	newUser := user.ToUser()

	if err := u.UserService.Create(newUser); err != nil {
		return err
	}

	ctx.JSON(200, newUser)
	return nil

}

func (u *UserController) Update(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	user, err := u.UserService.FindOne(uid)

	if err != nil {
		return err
	}

	var updateUser dtos.UpdateUser

	if err := ctx.Decode(&updateUser); err != nil {
		return err
	}
	updateUser.Decode(user)

	if err = u.UserService.Update(user); err != nil {
		return err
	}

	ctx.JSON(200, user)
	return nil
}

func (u *UserController) Delete(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}
	err = u.UserService.Delete(uid)

	if err == nil {
		ctx.JSON(200, struct {
			Message string `json:"message"`
		}{
			Message: "User Deleted",
		})
	}

	return err

}

func (u *UserController) Prefix() string {
	return "/api/v1/user"
}

func (u *UserController) Middlewares() []common.MiddleWare {
	return []common.MiddleWare{
		thor.LoggingMiddleware,
		thor.ErrorMiddleware,
		thor.PanicMiddleWare,
	}
}
