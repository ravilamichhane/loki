package authentication

import (
	"auth/iam/authentication/dto"
	"nest/common"
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
		common.POST("/login", u.Login),
		common.POST("/register", u.Register),
	}
}

func (u *AuthenticationController) Login(ctx common.HttpContext) error {

	return nil
}

func (u *AuthenticationController) Register(ctx common.HttpContext) error {
	var signUpRequest dto.SignUpRequest

	if err := ctx.Decode(&signUpRequest); err != nil {
		return err
	}

	user, err := u.AuthenticationService.Register(signUpRequest)

	if err != nil {
		return err
	}

	ctx.JSON(200, user)

	return nil
}

func (u *AuthenticationController) Prefix() string {

	return "/api/v1/auth"
}
