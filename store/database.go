package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/maha-1030/go-blog-api/models"
)

var (
	db *gorm.DB // db is used to store the database connection
)

// opens the databse connection
func init() {
	var err error

	if db, err = gorm.Open("sqlite3", "blog.db"); err != nil {
		panic("failed to connect database with error: " + err.Error())
	}

	fmt.Println("successfully connected to the database")

	db.AutoMigrate(&models.Tag{}, &models.Comment{}, &models.Post{}, &models.User{})

	fmt.Println("successfully completed migrations")
}
