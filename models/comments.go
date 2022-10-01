package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	AuthorID uint   `gorm:"author_id"`
	PostID   uint   `gorm:"post_id"`
	Message  string `gorm:"message"`
}
