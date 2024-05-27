package entities

import (
	"loki/common"
)

type User struct {
	common.Model
	FirstName string `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"-"`
}
