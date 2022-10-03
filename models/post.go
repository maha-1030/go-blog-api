package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	AuthorID uint      `gorm:"author_id"`                       // author of the post
	Title    string    `gorm:"title"`                           // title of the post
	Content  string    `gorm:"content"`                         // content in the post
	Tags     []Tag     `gorm:"many2many:post_tags;"`            // tags of the post
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"` // comments on the post
}

// Validate will validates the post data
func (p *Post) Validate() error {
	if p.AuthorID == 0 {
		return fmt.Errorf("missing field 'authorID'")
	}

	return nil
}

// BeforeSave is hook func that executes before saving post to validate data
func (p *Post) BeforeSave() error {
	return p.Validate()
}
