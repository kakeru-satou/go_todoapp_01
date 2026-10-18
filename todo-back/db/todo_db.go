package db

import (
	"go-learning/todoApp/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&models.Todo{}, &models.User{})
	if err != nil {
		panic(err)
	}
}
