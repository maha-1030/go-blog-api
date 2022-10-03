package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/service"
)

type comment struct {
	cs service.CommentService
}

// CommentHandlers is an object of CrudHandlers used to store comment handlers
var commentHandlers CrudHandlers

// NewCommentHandlers will create and return the new CommentHandlers
func NewCommentHandlers() CrudHandlers {
	return &comment{
		cs: service.GetCommentService(),
	}
}

// GetCommentHandlers will return the existing commentHandlers object or will create and return the new commentHandlers
func GetCommentHandlers() CrudHandlers {
	if commentHandlers == nil {
		commentHandlers = NewCommentHandlers()
	}

	return commentHandlers
}

// Create validates the comment post request and calls service layer to create comment
func (c *comment) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID, ok := vars["postID"]
	if !ok {
		fmt.Println("Missing path param 'postID' in the comment get request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'postID'")

		return
	}

	var commentRequest models.Comment

	if err := json.NewDecoder(r.Body).Decode(&commentRequest); err != nil {
		fmt.Println("Error while decoding the request body in comment create request, err: ", err)
		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into comment")

		return
	}

	username := r.Context().Value("username").(string)

	newComment, err := c.cs.Create(postID, username, &commentRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, newComment)
}

// Get validates the comment get request and calls service layer to get comment details
func (c *comment) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the comment get request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	existingComment, err := c.cs.Get(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, existingComment)
}

// Update validates the comment update request and calls service layer to update comment details
func (c *comment) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the comment update request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	postID, ok := vars["postID"]
	if !ok {
		fmt.Println("Missing path param 'postID' in the comment get request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'postID'")

		return
	}

	var commentRequest models.Comment

	if err := json.NewDecoder(r.Body).Decode(&commentRequest); err != nil {
		fmt.Println("Error while decoding the request body in comment update request, err: ", err)

		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into comment")
	}

	username := r.Context().Value("username").(string)

	updatedComment, err := c.cs.Update(id, postID, username, &commentRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, updatedComment)
}

// Delete validates the comment delete request and calls service layer to delete comment details
func (c *comment) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the comment delete request")

		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	username := r.Context().Value("username").(string)

	if err := c.cs.Delete(id, username); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
}
