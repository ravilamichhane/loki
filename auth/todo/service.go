package todo

import (
	"auth/todo/entities"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type TodoService interface {
	FindAll() ([]entities.Todo, error)
	FindOne(id uuid.UUID) (*entities.Todo, error)
	Create(Todo *entities.Todo) error
	Update(Todo *entities.Todo) error
	Delete(id uuid.UUID) error
}

type TodoServiceDB struct {
	db *gorm.DB
}

func NewTodoServiceDB(db *gorm.DB) TodoService {
	return &TodoServiceDB{db: db}

}

func (u *TodoServiceDB) FindOne(id uuid.UUID) (*entities.Todo, error) {
	todo := &entities.Todo{}
	if err := u.db.First(todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *TodoServiceDB) FindAll() ([]entities.Todo, error) {
	var todos []entities.Todo
	if err := u.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (u *TodoServiceDB) Create(todo *entities.Todo) error {
	return u.db.Create(&todo).Error
}

func (u *TodoServiceDB) Update(todo *entities.Todo) error {
	return u.db.Save(todo).Error
}

func (u *TodoServiceDB) Delete(id uuid.UUID) error {
	return u.db.Delete(&entities.Todo{}, id).Error
}
