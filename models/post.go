package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	AuthorID uint      `gorm:"author_id"`
	Title    string    `gorm:"title"`
	Content  string    `gorm:"content"`
	Tags     []Tag     `gorm:"many2many:post_tags;"`
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"`
}
