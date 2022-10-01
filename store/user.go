package store

import (
	"fmt"

	"github.com/maha-1030/go-blog-api/models"
	"gorm.io/gorm"
)

type user struct{}

// UserStore interface contains the User related methods that operate on database
type UserStore interface {
	Create(userRequest *models.User) (newUser *models.User, err error)
	Get(id int) (existingUser *models.User, err error)
	Update(id int, userRequest *models.User) (updatedUser *models.User, err error)
	Delete(id int) (err error)
}

// userStore is an object of UserStore
var userStore UserStore

// NewUserStore will create and return the new UserStore
func NewUserStore() UserStore {
	return &user{}
}

// GetUserStore will return the existing UserStore object or will create and return the new UserStore
func GetUserStore() UserStore {
	if userStore == nil {
		userStore = NewUserStore()
	}

	return userStore
}

// Create will create the new User in the database and responds with the newly created User and error if any
func (u *user) Create(userRequest *models.User) (newUser *models.User, err error) {
	if res := db.Create(userRequest); res.Error != nil {
		fmt.Println("Error while creating the new User, err: ", err)

		return nil, err
	}

	return userRequest, nil
}

// Get will retrieves the User with given UserID and responds with retrieved data and error if any
func (u *user) Get(id int) (existingUser *models.User, err error) {
	existingUser = &models.User{}

	if res := db.First(existingUser, id); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			fmt.Println("No User found with the ID: ", id)

			return nil, fmt.Errorf("no User found with the ID: %v", id)
		}

		fmt.Println("Error while retrieving a User with ID: ", id, ", err: ", err)

		return nil, err
	}

	return existingUser, nil
}

// Update will update the User with given UserID and responds with updated data and error if any
func (u *user) Update(id int, userRequest *models.User) (updatedUser *models.User, err error) {
	userRequest.ID = uint(id)
	if res := db.Save(userRequest); res.Error != nil {
		fmt.Println("Error while updating the User with ID: ", id, ", err: ", err)

		return nil, err
	}

	if updatedUser, err = u.Get(id); err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// Delete will delete User with given UserID and responds with error if any
func (u *user) Delete(id int) (err error) {
	if res := db.Delete(&models.User{}, id); res.Error != nil {
		fmt.Println("Error while deleting the User with ID: ", id, ", err: ", err)

		return err
	}

	return nil
}
