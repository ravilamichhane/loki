package aichat

import (
	"aichat/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type AichatService interface {
	FindAll() ([]entities.Aichat, error)
	FindOne(id uuid.UUID) (*entities.Aichat, error)
	Create(Aichat *entities.Aichat) error
	Update(Aichat *entities.Aichat) error
	Delete(id uuid.UUID) error
}

type AichatServiceDB struct {
	db *gorm.DB
}

func NewAichatServiceDB(db *gorm.DB) AichatService {
	return &AichatServiceDB{db: db}

}

func (u *AichatServiceDB) FindOne(id uuid.UUID) (*entities.Aichat, error) {
	aichat := &entities.Aichat{}
	if err := u.db.First(aichat, id).Error; err != nil {
		return nil, err
	}
	return aichat, nil
}

func (u *AichatServiceDB) FindAll() ([]entities.Aichat, error) {
	var aichats []entities.Aichat
	if err := u.db.Find(&aichats).Error; err != nil {
		return nil, err
	}
	return aichats, nil
}

func (u *AichatServiceDB) Create(aichat *entities.Aichat) error {
	return u.db.Create(&aichat).Error
}

func (u *AichatServiceDB) Update(aichat *entities.Aichat) error {
	return u.db.Save(aichat).Error
}

func (u *AichatServiceDB) Delete(id uuid.UUID) error {
	return u.db.Delete(&entities.Aichat{}, id).Error
}
