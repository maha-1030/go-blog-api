package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	TagLine string `gorm:"tag_line"`
}
