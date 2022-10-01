package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	AuthorID uint   `gorm:"author_id"` // author of the comment
	PostID   uint   `gorm:"post_id"`   // post on which comment is added
	Message  string `gorm:"message"`   // message in the comment
}
