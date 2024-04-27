package dto

import "auth/user/entities"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
	User         *entities.User `json:"user"`
}
