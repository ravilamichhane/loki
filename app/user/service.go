package user

import (
	"app/user/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser(id uuid.UUID) (*entities.User, error)
	GetUsers() ([]entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id uuid.UUID) error
}

type UserServiceDB struct {
	db *gorm.DB
}

func NewUserServiceDB(db *gorm.DB) UserService {
	return &UserServiceDB{db: db}

}

func (u *UserServiceDB) GetUser(id uuid.UUID) (*entities.User, error) {
	user := &entities.User{}
	if err := u.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceDB) GetUsers() ([]entities.User, error) {
	var users []entities.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserServiceDB) CreateUser(user *entities.User) error {
	return u.db.Create(&user).Error
}

func (u *UserServiceDB) UpdateUser(user *entities.User) error {
	return u.db.Save(user).Error
}

func (u *UserServiceDB) DeleteUser(id uuid.UUID) error {
	return u.db.Delete(&entities.User{}, id).Error
}
