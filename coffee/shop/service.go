package shop

import (
	"coffee/shop/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ShopService interface {
	FindAll() ([]entities.Shop, error)
	FindOne(id uuid.UUID) (*entities.Shop, error)
	Create(Shop *entities.Shop) error
	Update(Shop *entities.Shop) error
	Delete(id uuid.UUID) error
}

type ShopServiceDB struct {
	db *gorm.DB
}

func NewShopServiceDB(db *gorm.DB) ShopService {
	return &ShopServiceDB{db: db}

}

func (u *ShopServiceDB) FindOne(id uuid.UUID) (*entities.Shop, error) {
	shop := &entities.Shop{}
	if err := u.db.First(shop, id).Error; err != nil {
		return nil, err
	}
	return shop, nil
}

func (u *ShopServiceDB) FindAll() ([]entities.Shop, error) {
	var shops []entities.Shop
	if err := u.db.Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func (u *ShopServiceDB) Create(shop *entities.Shop) error {
	return u.db.Create(&shop).Error
}

func (u *ShopServiceDB) Update(shop *entities.Shop) error {
	return u.db.Save(shop).Error
}

func (u *ShopServiceDB) Delete(id uuid.UUID) error {
	return u.db.Delete(&entities.Shop{}, id).Error
}
