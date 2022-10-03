package service

import (
	"fmt"

	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/store"
)

type post struct {
	ps store.PostStore
	ts store.TagStore
	us store.UserStore
}

// PostService interface contains the Post related methods which have business logic
type PostService interface {
	Create(username string, postRequest *models.Post) (newPost *models.Post, err error)
	Get(idString string) (existingPost *models.Post, err error)
	Update(idString, username string, postRequest *models.Post) (updatedPost *models.Post, err error)
	Delete(idString, username string) (err error)
}

// postService is an object of PostService
var postService PostService

// NewPostService will create and return the new PostService
func NewPostService() PostService {
	return &post{
		ps: store.GetPostStore(),
		us: store.GetUserStore(),
	}
}

// GetPostService will return the existing PostService object or will create and return the new PostService
func GetPostService() PostService {
	if postService == nil {
		postService = NewPostService()
	}

	return postService
}

// Create updates post request with authorID & tagIDs and calls store layer to create the new Post
func (p *post) Create(username string, postRequest *models.Post) (newPost *models.Post, err error) {
	if postRequest == nil {
		return nil, fmt.Errorf("postDetails are not found")
	}

	existingUser, err := p.us.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	postRequest.AuthorID = existingUser.ID

	for i := range postRequest.Tags {
		t, err := p.ts.GetByTagLine(postRequest.Tags[i].TagLine)
		if err != nil {
			return nil, err
		}

		postRequest.Tags[i].ID = t.ID
	}

	return p.ps.Create(postRequest)
}

// Get validates the idString and calls store layer to get Post by ID
func (p *post) Get(idString string) (existingPost *models.Post, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	return p.ps.Get(id)
}

// Update checks for the existence of post and calls store layer to update requested fields of post
func (p *post) Update(idString, username string, postRequest *models.Post) (updatedPost *models.Post, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	existingUser, err := p.us.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	existingPost, err := p.ps.Get(id)
	if err != nil {
		return nil, err
	}

	if existingUser.ID != existingPost.AuthorID {
		return nil, fmt.Errorf("you are not allowed to update post of other user")
	}

	if postRequest.Content != "" {
		existingPost.Content = postRequest.Content
	}

	if postRequest.Title != "" {
		existingPost.Title = postRequest.Title
	}

	return p.ps.Update(id, existingPost)
}

// Delete will deletes the Post with given id
func (p *post) Delete(idString, username string) (err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return err
	}

	existingUser, err := p.us.GetByUsername(username)
	if err != nil {
		return err
	}

	existingPost, err := p.ps.Get(id)
	if err != nil {
		return nil
	}

	if existingUser.ID != existingPost.AuthorID {
		return fmt.Errorf("you are not allowed to delete post of other user")
	}

	return p.ps.Delete(id)
}
