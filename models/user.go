package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/maha-1030/go-blog-api/auth"
)

type User struct {
	gorm.Model
	Username         string    `gorm:"username"`                          // unique username for the user
	Password         string    `gorm:"password"`                          // password to authenticate the user
	FirstName        string    `gorm:"first_name"`                        // first name of the user
	LastName         string    `gorm:"last_name"`                         // last name of the user
	Age              int       `gorm:"age"`                               // age of the user
	Posts            []Post    `gorm:"foreignKey:AuthorID;references:ID"` // posts created by the user
	Comments         []Comment `gorm:"foreignKey:AuthorID;references:ID"` // comments added by the user
	isPasswordHashed bool      `gorm:"-" json:"-"`                        // used for marking when passwords hashed
}

// SetIsPasswordHashed will set the value of isPasswordHashed Field
func (u *User) SetIsPasswordHashed(isPasswordHashed bool) {
	u.isPasswordHashed = isPasswordHashed
}

// Validate will validates the user data
func (u *User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("missing field 'username'")
	}

	if u.Password == "" {
		return fmt.Errorf("missing field 'password'")
	}

	if len(u.Password) < 8 {
		return fmt.Errorf("password should have atleast 8 characters")
	}

	if u.FirstName == "" {
		return fmt.Errorf("missing field 'firstName'")
	}

	if u.LastName == "" {
		return fmt.Errorf("missing field 'lastName'")
	}

	if u.Age < 5 {
		return fmt.Errorf("only people of age 5 and above can become our users")
	}

	return nil
}

// BeforeSave is hook func that executes before saving user data to validate data and to hash password
func (u *User) BeforeSave() (err error) {
	if err = u.Validate(); err != nil {
		return err
	}

	if u.isPasswordHashed {
		return nil
	}

	u.Password, err = auth.Hash(u.Password)
	if err != nil {
		fmt.Println("unable to hash password, err: ", err.Error())
		return fmt.Errorf("internal server error")
	}

	return nil
}
