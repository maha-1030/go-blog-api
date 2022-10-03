package service

import (
	"fmt"

	"github.com/maha-1030/go-blog-api/models"
	"github.com/maha-1030/go-blog-api/store"
)

type tag struct {
	ts store.TagStore
}

// TagService interface contains the Tag related methods which have business logic
type TagService interface {
	Create(tagRequest *models.Tag) (newTag *models.Tag, err error)
	Get(idString string) (existingTag *models.Tag, err error)
	Update(idString string, tagRequest *models.Tag) (updatedTag *models.Tag, err error)
	Delete(idString string) (err error)
}

// tagService is an object of TagService
var tagService TagService

// NewTagService will create and return the new TagService
func NewTagService() TagService {
	return &tag{
		ts: store.GetTagStore(),
	}
}

// GetTagService will return the existing TagService object or will create and return the new TagService
func GetTagService() TagService {
	if tagService == nil {
		tagService = NewTagService()
	}

	return tagService
}

// Create validates the tagline and calls store layer to create the new Tag
func (t *tag) Create(tagRequest *models.Tag) (newTag *models.Tag, err error) {
	if tagRequest == nil {
		return nil, fmt.Errorf("tagDetails are not found")
	}

	if existingTag, _ := t.ts.GetByTagLine(tagRequest.TagLine); existingTag != nil {
		return nil, fmt.Errorf("tagname is already in use")
	}

	return t.ts.Create(tagRequest)
}

// Get validates the idString and calls store layer to get Tag by ID
func (t *tag) Get(idString string) (existingTag *models.Tag, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	if existingTag, err = t.ts.Get(id); err != nil {
		return nil, err
	}

	return
}

// Update checks for the existence of tag and calls store layer to update the tag
func (t *tag) Update(idString string, tagRequest *models.Tag) (updatedTag *models.Tag, err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return nil, err
	}

	_, err = t.ts.Get(id)
	if err != nil {
		return nil, err
	}

	if tagWithTagline, _ := t.ts.GetByTagLine(tagRequest.TagLine); tagWithTagline != nil {
		return nil, fmt.Errorf("this tagline is already in use")
	}

	return t.ts.Update(id, tagRequest)
}

// Delete will deletes the Tag with given id
func (t *tag) Delete(idString string) (err error) {
	id, err := getIDFromString(idString)
	if err != nil {
		return err
	}

	_, err = t.ts.Get(id)
	if err != nil {
		return err
	}

	return t.ts.Delete(id)
}
