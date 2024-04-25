package item

import (
	"coffee/item/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ItemService interface {
	FindAll() ([]entities.Item, error)
	FindOne(id uuid.UUID) (*entities.Item, error)
	Create(Item *entities.Item) error
	Update(Item *entities.Item) error
	Delete(id uuid.UUID) error
}

type ItemServiceDB struct {
	db *gorm.DB
}

func NewItemServiceDB(db *gorm.DB) ItemService {
	return &ItemServiceDB{db: db}

}

func (u *ItemServiceDB) FindOne(id uuid.UUID) (*entities.Item, error) {
	item := &entities.Item{}
	if err := u.db.First(item, id).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (u *ItemServiceDB) FindAll() ([]entities.Item, error) {
	var items []entities.Item
	if err := u.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (u *ItemServiceDB) Create(item *entities.Item) error {
	return u.db.Create(&item).Error
}

func (u *ItemServiceDB) Update(item *entities.Item) error {
	return u.db.Save(item).Error
}

func (u *ItemServiceDB) Delete(id uuid.UUID) error {
	return u.db.Delete(&entities.Item{}, id).Error
}
