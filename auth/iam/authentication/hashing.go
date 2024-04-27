package authentication

import "golang.org/x/crypto/bcrypt"

type HashingService interface {
	Hash(data string) (string, error)
	Compare(hashed, password string) bool
}

type BcryptService struct {
}

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

func (b BcryptService) Hash(data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (b BcryptService) Compare(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
