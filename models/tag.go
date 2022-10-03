package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	TagLine string `gorm:"tag_line"` // tag line used to represent the tag
}

// Validate will validates the tag
func (t *Tag) Validate() error {
	if t.TagLine == "" {
		return fmt.Errorf("missing field 'tagline'")
	}

	return nil
}

// BeforeSave is hook func that executes before saving tag to validate data
func (t *Tag) BeforeSave() error {
	return t.Validate()
}
