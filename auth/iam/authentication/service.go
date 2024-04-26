package authentication

import (
	"auth/iam/authentication/dto"
	"auth/user"
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

func (u *AuthenticationService) Register(signUpRequest dto.SignUpRequest) (dto.SignUpResponse, error) {

	createUserRequest := signUpRequest.ToCreateUser()

	user := createUserRequest.ToUser()

	if err := u.userService.Create(user); err != nil {
		return dto.SignUpResponse{}, err
	}

	return dto.SignUpResponse{
		Token:        "token",
		RefreshToken: "refresh",
		User:         user,
	}, nil
}
