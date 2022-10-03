package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/service"
)

type user struct {
	us service.UserService
}

// authHandlers is an object of AuthHandlers used to store authentication handlers
var authHandlers AuthHandlers

// NewAuthHandlers will create and return the new UserHandlers
func NewAuthHandlers() AuthHandlers {
	return &user{
		us: service.GetUserService(),
	}
}

// GetAuthHandlers will return the existing AuthHandlers object or will create and return the new AuthHandlers
func GetAuthHandlers() AuthHandlers {
	if authHandlers == nil {
		authHandlers = NewAuthHandlers()
	}

	return authHandlers
}

// userHandlers is an object of CrudHandlers used to store user handlers
var userHandlers CrudHandlers

// NewUserHandlers will create and return the new UserHandlers
func NewUserHandlers() CrudHandlers {
	return &user{
		us: service.GetUserService(),
	}
}

// GetUserHandlers will return the existing UserHandlers object or will create and return the new UserHandlers
func GetUserHandlers() CrudHandlers {
	if userHandlers == nil {
		userHandlers = NewUserHandlers()
	}

	return userHandlers
}

// Register will validates the user post request and registers an user by calling service layer
func (u *user) Register(w http.ResponseWriter, r *http.Request) {
	u.Create(w, r)
}

// Login will validates the user login request and authenticates the user by calling service layer
func (u *user) Login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	token, err := u.us.Login(username, password)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"token": *token})
}

// Create validates the user post request and calls service layer to create user
func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		fmt.Println("Error while decoding the request body in user create request, err: ", err)

		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into user")

		return
	}

	newUser, err := u.us.Create(&userRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, newUser)
}

// Get validates the user get request and calls service layer to get user details
func (u *user) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the user get request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	username := r.Context().Value("username").(string)

	existingUser, err := u.us.Get(id, username)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, existingUser)
}

// Update validates the user update request and calls service layer to update user details
func (u *user) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the user update request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	var userRequest models.User

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		fmt.Println("Error while decoding the request body in user update request, err: ", err)

		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into user")
	}

	username := r.Context().Value("username").(string)

	updatedUser, err := u.us.Update(id, username, &userRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, updatedUser)
}

// Delete validates the user delete request and calls service layer to delete user details
func (u *user) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the user delete request")

		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	username := r.Context().Value("username").(string)

	if err := u.us.Delete(id, username); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
}
