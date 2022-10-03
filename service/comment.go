package service

import (
	"fmt"

	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/store"
)

type comment struct {
	cs store.CommentStore
	ps store.PostStore
	us store.UserStore
}

// CommentService interface contains the Comment related methods which have business logic
type CommentService interface {
	Create(postIdString, username string, commentRequest *models.Comment) (newComment *models.Comment, err error)
	Get(idString string) (existingComment *models.Comment, err error)
	Update(idString, postIdString, username string, commentRequest *models.Comment) (updatedComment *models.Comment, err error)
	Delete(idString, username string) (err error)
}

// commentService is an object of CommentService
var commentService CommentService

// NewCommentService will create and return the new CommentService
func NewCommentService() CommentService {
	return &comment{
		cs: store.GetCommentStore(),
		ps: store.GetPostStore(),
		us: store.GetUserStore(),
	}
}

// GetCommentService will return the existing CommentService object or will create and return the new CommentService
func GetCommentService() CommentService {
	if commentService == nil {
		commentService = NewCommentService()
	}

	return commentService
}

// Create checks for user & post and calls store layer to create the new Comment
func (c *comment) Create(postIdString, username string, commentRequest *models.Comment) (newComment *models.Comment, err error) {
	postID, err := getIDFromString(postIdString)
	if err != nil {
		return nil, err
	}

	if commentRequest == nil {
		return nil, fmt.Errorf("comment details are not found")
	}

	existingUser, err := c.us.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, fmt.Errorf("no user found with the authorized username: %v", username)
	}

	existingPost, err := c.ps.Get(postID)
	if err != nil {
		return nil, err
	}
	if existingPost == nil {
		return nil, fmt.Errorf("no post found with the id: %v", postID)
	}

	commentRequest.AuthorID = existingUser.ID
	commentRequest.PostID = existingPost.ID

	return c.cs.Create(commentRequest)
}

// Get validates the idString and calls store layer to get Comment by ID
func (c *comment) Get(idString string) (existingComment *models.Comment, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	if existingComment, err = c.cs.Get(id); err != nil {
		return nil, err
	}
	if existingComment == nil {
		return nil, fmt.Errorf("no comment found with id: %v", id)
	}

	return existingComment, nil
}

// Update checks for the existence of comment and calls store layer to update the Comment
func (c *comment) Update(idString, postIdString, username string, commentRequest *models.Comment) (updatedComment *models.Comment, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	postID, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	existingUser, err := c.us.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, fmt.Errorf("no user found with the authorized username: %v", username)
	}

	existingPost, err := c.ps.Get(postID)
	if err != nil {
		return nil, err
	}
	if existingPost == nil {
		return nil, fmt.Errorf("no post found with the id: %v", postID)
	}

	existingComment, err := c.cs.Get(id)
	if err != nil {
		return nil, err
	}
	if existingComment == nil {
		return nil, fmt.Errorf("no comment found with id: %v", id)
	}

	if existingUser.ID != existingComment.AuthorID {
		return nil, fmt.Errorf("you are not allowed to update comment of other user")
	}

	if existingPost.ID != existingComment.PostID {
		return nil, fmt.Errorf("comment doesn't belongs to the post provided in request")
	}

	existingComment.Message = commentRequest.Message

	return c.cs.Update(id, existingComment)
}

// Delete will deletes the Comment with given id
func (c *comment) Delete(idString, username string) (err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return err
	}

	existingUser, err := c.us.GetByUsername(username)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return fmt.Errorf("no user found with the authorized username: %v", username)
	}

	existingComment, err := c.cs.Get(id)
	if err != nil {
		return nil
	}
	if existingComment == nil {
		return fmt.Errorf("no comment found with id: %v", id)
	}

	if existingUser.ID != existingComment.AuthorID {
		return fmt.Errorf("you are not allowed to delete comment of other user")
	}

	return c.cs.Delete(id)
}
