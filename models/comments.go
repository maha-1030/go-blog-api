package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	AuthorID uint   `gorm:"author_id"` // author of the comment
	PostID   uint   `gorm:"post_id"`   // post on which comment is added
	Message  string `gorm:"message"`   // message in the comment
}

// Validate will validates the comment
func (c *Comment) Validate() error {
	if c.AuthorID == 0 {
		return fmt.Errorf("missing field 'authorID'")
	}

	if c.PostID == 0 {
		return fmt.Errorf("missing field 'postID'")
	}

	return nil
}

// BeforeSave is hook func that executes before saving post to validate data
func (c *Comment) BeforeSave() error {
	return c.Validate()
}
