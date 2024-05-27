package authentication

import (
	"auth/iam/authentication/dto"
	"log"
	"loki/common"
	"loki/thor"
)

type AuthenticationController struct {
	AuthenticationService AuthenticationService
}

func NewAuthenticationController(AuthenticationService AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		AuthenticationService: AuthenticationService,
	}
}

func (u *AuthenticationController) Routes() []common.Route {
	return []common.Route{
		common.POST("/signin", u.SignIn),
		common.POST("/signup", u.SignUp),
	}
}

func (u *AuthenticationController) SignIn(ctx common.HttpContext) error {

	return nil
}

func (u *AuthenticationController) SignUp(ctx common.HttpContext) error {
	var signUpRequest dto.SignUpRequest

	if err := ctx.Decode(&signUpRequest); err != nil {
		return err
	}

	user, err := u.AuthenticationService.Register(signUpRequest)

	log.Println("User: ", user)

	if err != nil {
		return err
	}

	ctx.JSON(201, common.NewSuccessResponse("User created successfully", user))

	return nil
}

func (u *AuthenticationController) Prefix() string {

	return "/api/v1/auth"
}

func (u *AuthenticationController) Middlewares() []common.MiddleWare {
	return []common.MiddleWare{
		thor.LoggingMiddleware,
		thor.ErrorMiddleware,
		thor.PanicMiddleWare,
		thor.CORSMiddleWare,
	}
}
