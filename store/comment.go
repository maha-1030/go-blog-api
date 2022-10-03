package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/maha-1030/go-blog-api/models"
)

type comment struct{}

// CommentStore interface contains the Comment related methods that operate on database
type CommentStore interface {
	Create(commentRequest *models.Comment) (newComment *models.Comment, err error)
	Get(id int) (existingComment *models.Comment, err error)
	Update(id int, commentRequest *models.Comment) (updatedComment *models.Comment, err error)
	Delete(id int) (err error)
}

// commentStore is an object of CommentStore
var commentStore CommentStore

// NewCommentStore will create and return the new CommentStore
func NewCommentStore() CommentStore {
	return &comment{}
}

// GetCommentStore will return the existing CommentStore object or will create and return the new CommentStore
func GetCommentStore() CommentStore {
	if commentStore == nil {
		commentStore = NewCommentStore()
	}

	return commentStore
}

// Create will create the new Comment in the database and responds with the newly created Comment and error if any
func (c *comment) Create(commentRequest *models.Comment) (newComment *models.Comment, err error) {
	if err := db.Create(commentRequest).Error; err != nil {
		fmt.Println("Error while creating the new Comment, err: ", err)

		return nil, err
	}

	return commentRequest, nil
}

// Get will retrieves the Comment with given CommentID and responds with retrieved data and error if any
func (c *comment) Get(id int) (existingComment *models.Comment, err error) {
	existingComment = &models.Comment{}

	if err := db.First(existingComment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		fmt.Println("Error while retrieving a Comment with ID: ", id, ", err: ", err)

		return nil, err
	}

	return existingComment, nil
}

// Update will update the Comment with given CommentID and responds with updated data and error if any
func (c *comment) Update(id int, commentRequest *models.Comment) (updatedComment *models.Comment, err error) {
	commentRequest.ID = uint(id)
	if err := db.Save(commentRequest).Error; err != nil {
		fmt.Println("Error while updating the Comment with ID: ", id, ", err: ", err)

		return nil, err
	}

	return c.Get(id)
}

// Delete will delete Comment with given CommentID and responds with error if any
func (c *comment) Delete(id int) (err error) {
	if err := db.Delete(&models.Comment{}, id).Error; err != nil {
		fmt.Println("Error while deleting the Comment with ID: ", id, ", err: ", err)

		return err
	}

	return nil
}
