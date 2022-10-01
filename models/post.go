package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	AuthorID uint      `gorm:"author_id"`                       // author of the post
	Title    string    `gorm:"title"`                           // title of the post
	Content  string    `gorm:"content"`                         // content in the post
	Tags     []Tag     `gorm:"many2many:post_tags;"`            // tags of the post
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"` // comments on the post
}
