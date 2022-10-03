package service

import (
	"fmt"

	"github.com/maha-1030/go-blog-api/auth"
	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/store"
)

type user struct {
	us store.UserStore
}

// UserService interface contains the User related methods which have business logic
type UserService interface {
	Login(username, password string) (token *string, err error)
	Create(userRequest *models.User) (newUser *models.User, err error)
	Get(idString, username string) (existingUser *models.User, err error)
	Update(idString, username string, userRequest *models.User) (updatedUser *models.User, err error)
	Delete(idString, username string) (err error)
}

// userService is an object of UserService
var userService UserService

// NewUserService will create and return the new UserService
func NewUserService() UserService {
	return &user{
		us: store.GetUserStore(),
	}
}

// GetUserService will return the existing UserService object or will create and return the new UserService
func GetUserService() UserService {
	if userService == nil {
		userService = NewUserService()
	}

	return userService
}

// Login validates user credentials and responds with jwt token & error if any
func (u *user) Login(username, password string) (token *string, err error) {
	existingUser, err := u.us.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if err = auth.VerifyPassword(existingUser.Password, password); err != nil {
		fmt.Println("error occured while verifying password, err: ", err)

		return nil, fmt.Errorf("incorrect password")
	}

	if token, err = auth.CreateToken(username); err != nil {
		fmt.Println("unable to create token, err: ", err)

		return nil, fmt.Errorf("unable to create token, err: %v", err)
	}

	return token, nil
}

// Create validates the Username and calls store layer to create the new User
func (u *user) Create(userRequest *models.User) (newUser *models.User, err error) {
	if userRequest == nil {
		return nil, fmt.Errorf("userDetails are not found")
	}

	if existingUser, _ := u.us.GetByUsername(userRequest.Username); existingUser != nil {
		return nil, fmt.Errorf("username is already in use")
	}

	return u.us.Create(userRequest)
}

// Get validates the idString and calls store layer to get User by ID
func (u *user) Get(idString, username string) (existingUser *models.User, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	if existingUser, err = u.us.Get(id); err != nil {
		return nil, err
	}

	if existingUser.Username != username {
		return nil, fmt.Errorf("you are not allowed to get other user details")
	}

	return
}

// Update checks for the existence of user and calls store layer to update requested fields of user
func (u *user) Update(idString, username string, userRequest *models.User) (updatedUser *models.User, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	existingUser, err := u.us.Get(id)
	if err != nil {
		return nil, err
	}

	if existingUser.Username != username {
		return nil, fmt.Errorf("you are not allowed to update other user details")
	}

	if userRequest.Age == 0 {
		userRequest.Age = existingUser.Age
	}

	if userRequest.FirstName == "" {
		userRequest.FirstName = existingUser.FirstName
	}

	if userRequest.LastName == "" {
		userRequest.LastName = existingUser.LastName
	}

	if userRequest.Password == "" {
		userRequest.Password = existingUser.Password
		userRequest.SetIsPasswordHashed(true)
	}

	if userRequest.Username == "" {
		userRequest.Username = existingUser.Username
	} else if userRequest.Username != existingUser.Username {
		if userWithNewUsername, _ := u.us.GetByUsername(userRequest.Username); userWithNewUsername != nil {
			return nil, fmt.Errorf("username is already in use")
		}
	}

	return u.us.Update(id, userRequest)
}

// Delete will deletes the User with given id
func (u *user) Delete(idString, username string) (err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return err
	}

	existingUser, err := u.us.Get(id)
	if err != nil {
		return err
	}

	if existingUser.Username != username {
		return fmt.Errorf("you are not allowed to delete other user details")
	}

	return u.us.Delete(id)
}
