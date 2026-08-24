package db

import (
	"go-learning/todoApp/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	var err error

	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Todo{}, &models.User{})
}
