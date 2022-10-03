package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/maha-1030/go-blog-api/models"
)

type post struct{}

// PostStore interface contains the Post related methods that operate on database
type PostStore interface {
	Create(postRequest *models.Post) (newPost *models.Post, err error)
	Get(id int) (existingPost *models.Post, err error)
	Update(id int, postRequest *models.Post) (updatedPost *models.Post, err error)
	Delete(id int) (err error)
}

// postStore is an object of PostStore
var postStore PostStore

// NewPostStore will create and return the new PostStore
func NewPostStore() PostStore {
	return &post{}
}

// GetPostStore will return the existing PostStore object or will create and return the new PostStore
func GetPostStore() PostStore {
	if postStore == nil {
		postStore = NewPostStore()
	}

	return postStore
}

// Create will create the new Post in the database and responds with the newly created Post and error if any
func (p *post) Create(postRequest *models.Post) (newPost *models.Post, err error) {
	if err := db.Create(postRequest).Error; err != nil {
		fmt.Println("Error while creating the new Post, err: ", err)

		return nil, err
	}

	return postRequest, nil
}

// Get will retrieves the Post with given PostID and responds with retrieved data and error if any
func (p *post) Get(id int) (existingPost *models.Post, err error) {
	existingPost = &models.Post{}

	if err := db.First(existingPost, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		fmt.Println("Error while retrieving a Post with ID: ", id, ", err: ", err)

		return nil, err
	}

	return existingPost, nil
}

// Update will update the Post with given PostID and responds with updated data and error if any
func (p *post) Update(id int, postRequest *models.Post) (updatedPost *models.Post, err error) {
	postRequest.ID = uint(id)
	if err := db.Save(postRequest).Error; err != nil {
		fmt.Println("Error while updating the Post with ID: ", id, ", err: ", err)

		return nil, err
	}

	return p.Get(id)
}

// Delete will delete Post with given PostID and responds with error if any
func (p *post) Delete(id int) (err error) {
	if err := db.Delete(&models.Post{}, id).Error; err != nil {
		fmt.Println("Error while deleting the Post with ID: ", id, ", err: ", err)

		return err
	}

	return nil
}
