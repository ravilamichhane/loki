package user

import (
	"app/dtos"
	"nest/common"
	"nest/thor"

	"github.com/google/uuid"
)

type UserController struct {
	UserService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		UserService: userService,
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

	users, err := u.UserService.GetUsers()

	if err != nil {
		return err
	}

	ctx.JSON(200, users)
	return nil
}

func (u *UserController) FindOne(ctx common.HttpContext) error {
	id := ctx.GetParam("id")
	uid, err := uuid.FromBytes([]byte(id))

	if err != nil {
		return thor.NewTrustedError("Invalid UUID", err, 400)
	}

	user, err := u.UserService.GetUser(uid)

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

	if err := u.UserService.CreateUser(user.ToUser()); err != nil {
		return err
	}

	ctx.JSON(200, user)
	return nil

}

func (u *UserController) Update(ctx common.HttpContext) error {
	ctx.Text(200, "User")
	return nil
}

func (u *UserController) Delete(ctx common.HttpContext) error {
	ctx.Text(200, "User")
	return nil
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
