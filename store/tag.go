package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/maha-1030/go-blog-api/models"
)

type tag struct{}

// TagStore interface contains the Tag related methods that operate on database
type TagStore interface {
	Create(tagRequest *models.Tag) (newTag *models.Tag, err error)
	Get(id int) (existingTag *models.Tag, err error)
	GetByTagLine(tagLine string) (existingTag *models.Tag, err error)
	Update(id int, tagRequest *models.Tag) (updatedTag *models.Tag, err error)
	Delete(id int) (err error)
}

// tagStore is an object of TagStore
var tagStore TagStore

// NewTagStore will create and return the new TagStore
func NewTagStore() TagStore {
	return &tag{}
}

// GetTagStore will return the existing TagStore object or will create and return the new TagStore
func GetTagStore() TagStore {
	if tagStore == nil {
		tagStore = NewTagStore()
	}

	return tagStore
}

// Create will create the new Tag in the database and responds with the newly created Tag and error if any
func (t *tag) Create(tagRequest *models.Tag) (newTag *models.Tag, err error) {
	if res := db.Create(tagRequest); res.Error != nil {
		fmt.Println("Error while creating the new Tag, err: ", err)

		return nil, err
	}

	return tagRequest, nil
}

// Get will retrieves the Tag with given TagID and responds with retrieved data and error if any
func (t *tag) Get(id int) (existingTag *models.Tag, err error) {
	existingTag = &models.Tag{}

	if res := db.First(existingTag, id); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			fmt.Println("No Tag found with the ID: ", id)

			return nil, fmt.Errorf("no Tag found with the ID: %v", id)
		}

		fmt.Println("Error while retrieving a Tag with ID: ", id, ", err: ", err)

		return nil, err
	}

	return existingTag, nil
}

// GetByTagLine will retrieves the Tag with given tagline and responds with retrieved data and error if any
func (t *tag) GetByTagLine(tagLine string) (existingTag *models.Tag, err error) {
	existingTag = &models.Tag{}

	if res := db.Where("tag_line = ?", tagLine).First(existingTag); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			fmt.Println("No Tag found with the tagline: ", tagLine)

			return nil, fmt.Errorf("no Tag found with the tagline: %v", tagLine)
		}

		fmt.Println("Error while retrieving a Tag with tagline: ", tagLine, ", err: ", err)

		return nil, err
	}

	return existingTag, nil
}

// Update will update the Tag with given TagID and responds with updated data and error if any
func (t *tag) Update(id int, tagRequest *models.Tag) (updatedTag *models.Tag, err error) {
	tagRequest.ID = uint(id)
	if res := db.Save(tagRequest); res.Error != nil {
		fmt.Println("Error while updating the Tag with ID: ", id, ", err: ", err)

		return nil, err
	}

	if updatedTag, err = t.Get(id); err != nil {
		return nil, err
	}

	return updatedTag, nil
}

// Delete will delete Tag with given TagID and responds with error if any
func (t *tag) Delete(id int) (err error) {
	if res := db.Delete(&models.Tag{}, id); res.Error != nil {
		fmt.Println("Error while deleting the Tag with ID: ", id, ", err: ", err)

		return err
	}

	return nil
}
