package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/service"
)

type tag struct {
	ts service.TagService
}

// TagHandlers is an object of CrudHandlers used to store Tag handlers
var tagHandlers CrudHandlers

// NewTagHandlers will create and return the new TagHandlers
func NewTagHandlers() CrudHandlers {
	return &tag{
		ts: service.GetTagService(),
	}
}

// GetTagHandlers will return the existing TagHandlers object or will create and return the new TagHandlers
func GetTagHandlers() CrudHandlers {
	if tagHandlers == nil {
		tagHandlers = NewTagHandlers()
	}

	return tagHandlers
}

// Create validates the Tag post request and calls service layer to create tag
func (t *tag) Create(w http.ResponseWriter, r *http.Request) {
	var tagRequest models.Tag

	if err := json.NewDecoder(r.Body).Decode(&tagRequest); err != nil {
		fmt.Println("Error while decoding the request body in tag create request, err: ", err)

		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into tag")

		return
	}

	newTag, err := t.ts.Create(&tagRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, newTag)
}

// Get validates the tag get request and calls service layer to get tag details
func (t *tag) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the tag get request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	existingTag, err := t.ts.Get(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, existingTag)
}

// Update validates the tag update request and calls service layer to update tag details
func (t *tag) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the tag update request")
		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	var tagRequest models.Tag

	if err := json.NewDecoder(r.Body).Decode(&tagRequest); err != nil {
		fmt.Println("Error while decoding the request body in tag update request, err: ", err)

		RespondWithError(w, http.StatusBadRequest, "unable to decode request body into tag")
	}

	updatedTag, err := t.ts.Update(id, &tagRequest)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, updatedTag)
}

// Delete validates the tag delete request and calls service layer to delete tag details
func (t *tag) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Missing path param 'id' in the tag delete request")

		RespondWithError(w, http.StatusBadRequest, "missing path param 'id'")

		return
	}

	if err := t.ts.Delete(id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"status": "success"})
}
