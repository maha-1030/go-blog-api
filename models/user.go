package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"username"`                          // unique username for the user
	Password  string    `gorm:"password"`                          // password to authenticate the user
	FirstName string    `gorm:"first_name"`                        // first name of the user
	LastName  string    `gorm:"last_name"`                         // last name of the user
	Age       int       `gorm:"age"`                               // age of the user
	Posts     []Post    `gorm:"foreignKey:AuthorID;references:ID"` // posts created by the user
	Comments  []Comment `gorm:"foreignKey:AuthorID;references:ID"` // comments added by the user
}
