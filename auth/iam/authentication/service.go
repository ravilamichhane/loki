package authentication

import (
	"auth/iam/authentication/dto"
	"auth/user"
	"auth/user/entities"
)

type AuthenticationService struct {
	userService user.UserService
}

func NewAuthenticationService(userService user.UserService) AuthenticationService {
	return AuthenticationService{
		userService: userService,
	}
}

func (u *AuthenticationService) Authenticate(username string, password string) (bool, error) {
	return true, nil
}

func (u *AuthenticationService) Register(signUpRequest dto.SignUpRequest) (*entities.User, error) {

	
	return nil, nil
}
