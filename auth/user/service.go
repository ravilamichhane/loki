package user

import (
	"auth/user/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserService interface {
	FindAll() ([]entities.User, error)
	FindOne(id uuid.UUID) (*entities.User, error)
	Create(User *entities.User) error
	Update(User *entities.User) error
	Delete(id uuid.UUID) error
}

type UserServiceDB struct {
	db *gorm.DB
}

func NewUserServiceDB(db *gorm.DB) UserService {
	return &UserServiceDB{db: db}

}

func (u *UserServiceDB) FindOne(id uuid.UUID) (*entities.User, error) {
	user := &entities.User{}
	if err := u.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceDB) FindAll() ([]entities.User, error) {
	var users []entities.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserServiceDB) Create(user *entities.User) error {
	return u.db.Create(&user).Error
}

func (u *UserServiceDB) Update(user *entities.User) error {
	return u.db.Save(user).Error
}

func (u *UserServiceDB) Delete(id uuid.UUID) error {
	return u.db.Delete(&entities.User{}, id).Error
}
