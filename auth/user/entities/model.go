package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
}
