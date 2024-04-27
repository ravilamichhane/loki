package authentication

import (
	"auth/iam/authentication/dto"
	"auth/user"
)

type AuthenticationService struct {
	userService    user.UserService
	hasgingService HashingService
	jwtService     JwtService
}

func NewAuthenticationService(userService user.UserService) AuthenticationService {
	return AuthenticationService{
		userService:    userService,
		hasgingService: NewBcryptService(),
		jwtService:     NewJwtService(),
	}
}

func (u *AuthenticationService) Authenticate(signInRequest dto.SignInRequest) (dto.SignInRequest, error) {
	user, err := u.userService.FindOneByEmail(signInRequest.Email)

	if err != nil {
		return dto.SignInRequest{}, err
	}

	if !u.hasgingService.Compare(user.Password, signInRequest.Password) {
		return dto.SignInRequest{}, nil
	}

	return dto.SignInRequest{}, nil
}

func (u *AuthenticationService) Register(signUpRequest dto.SignUpRequest) (dto.SignUpResponse, error) {

	createUserRequest := signUpRequest.ToCreateUser()

	user := createUserRequest.ToUser()

	hashedPassword, err := u.hasgingService.Hash(user.Password)

	if err != nil {
		return dto.SignUpResponse{}, err
	}

	user.Password = hashedPassword

	if err := u.userService.Create(user); err != nil {
		return dto.SignUpResponse{}, err
	}

	tokens, err := u.jwtService.GenerateToken(JwtUser{
		ID:    user.ID,
		Email: user.Email,
	})

	if err != nil {
		return dto.SignUpResponse{}, err
	}

	return dto.SignUpResponse{
		Token:        tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         user,
		HasSetup:     false,
	}, nil
}
