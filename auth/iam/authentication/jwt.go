package authentication

import (
	"errors"
	"loki/thor"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtUser struct {
	ID    uuid.UUID `json:"user_id"`
	Role  string    `json:"role"`
	Email string    `json:"email"`
}

type Claims struct {
	jwt.RegisteredClaims
	Subject int    `json:"sub"`
	Role    string `json:"role"`
	Email   string `json:"email"`
}

type JwtService interface {
	GenerateToken(user JwtUser) (JwtResponse, error)
	ValidateToken(token string) (*Claims, error)
	ValidateAppToken(token string) (bool, string)
}

type DefaultJwtService struct {
	Issuer        string
	Audience      string
	SecretKey     string
	AppSecret     string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
}

func NewJwtService() JwtService {
	issuer := thor.GetEnv("JWT_ISSUER", "tetoteto")
	audience := thor.GetEnv("JWT_AUDIENCE", "tetoteto")
	secretKey := thor.GetEnv("JWT_SECRET", "tetoteto")
	appSecret := thor.GetEnv("APP_SECRET", "tetoteto")
	tokenExpiry := thor.GetEnvInt("TOKEN_EXPIRY", 20)
	refreshExpiry := thor.GetEnvInt("REFRESH_EXPIRY", 86400000)

	return &DefaultJwtService{
		Issuer:        issuer,
		Audience:      audience,
		SecretKey:     secretKey,
		AppSecret:     appSecret,
		TokenExpiry:   time.Duration(tokenExpiry) * time.Second,
		RefreshExpiry: time.Duration(refreshExpiry) * time.Second,
	}

}

func (s *DefaultJwtService) GenerateToken(user JwtUser) (JwtResponse, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["role"] = user.Role
	claims["email"] = user.Email
	claims["aud"] = s.Audience
	claims["iss"] = s.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"
	claims["exp"] = time.Now().UTC().Add(s.TokenExpiry).Unix()

	signedToken, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return JwtResponse{}, err
	}

	// Refresh token

	refreshToken := jwt.New(jwt.SigningMethodHS512)

	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = user.ID
	refreshClaims["iat"] = time.Now().UTC().Unix()
	refreshClaims["typ"] = "JWT"
	refreshClaims["exp"] = time.Now().UTC().Add(s.RefreshExpiry).Unix()

	signedRefreshToken, err := refreshToken.SignedString([]byte(s.SecretKey))
	if err != nil {
		return JwtResponse{}, err
	}

	return JwtResponse{
		AccessToken:  signedToken,
		RefreshToken: signedRefreshToken,
	}, nil

}

func (s *DefaultJwtService) ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})

	if err != nil {

		if err == jwt.ErrTokenExpired {
			return nil, thor.NewTrustedError(errors.New("token expired"), 401)
		}
		return nil, thor.NewTrustedError(err, 401)
	}
	return claims, nil
}

func (s *DefaultJwtService) ValidateAppToken(token string) (bool, string) {
	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.AppSecret), nil
	})

	if err != nil {
		return false, ""
	}

	sub, err := payload.Claims.GetSubject()

	if err != nil {
		return false, ""
	}

	return true, sub
}
