package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"username"`
	Password  string `gorm:"password"`
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
	Age       int    `gorm:"age"`
}
