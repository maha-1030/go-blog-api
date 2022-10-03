package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/service"
)

type post struct {
	ps service.PostService
}

// PostHandlers is an object of CrudHandlers used to store post handlers
var postHandlers CrudHandlers

// NewPostHandlers will create and return the new PostHandlers
func NewPostHandlers() CrudHandlers {
	return &post{
		ps: service.GetPostService(),
	}
}

// GetPostHandlers will return the existing PostHandlers object or will create and return the new PostHandlers
func GetPostHandlers() CrudHandlers {
	if postHandlers == nil {
		postHandlers = NewPostHandlers()
	}

	return postHandlers
}

// Create validates the post post request and calls service layer to create post
func (p *post) Create(w http.ResponseWriter, r *http.Request) {
	var postRequest models.Post

	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		fmt.Println("Error while decoding the request body in post create request, err: ", err)

		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into post")

		return
	}

	username := r.Context().Value("username").(string)

	newPost, err := p.ps.Create(username, &postRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, newPost)
}

// Get validates the post get request and calls service layer to get post details
func (p *post) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the post get request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	existingPost, err := p.ps.Get(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, existingPost)
}

// Update validates the post update request and calls service layer to update post details
func (p *post) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the post update request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	var postRequest models.Post

	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		fmt.Println("Error while decoding the request body in post update request, err: ", err)

		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into post")
	}

	username := r.Context().Value("username").(string)

	updatedPost, err := p.ps.Update(id, username, &postRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, updatedPost)
}

// Delete validates the post delete request and calls service layer to delete post details
func (p *post) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the post delete request")

		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	username := r.Context().Value("username").(string)

	if err := p.ps.Delete(id, username); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
}
