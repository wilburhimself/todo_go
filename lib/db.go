package lib

import (
	"github.com/wilburhimself/go_todo_list/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Todo{})
}

func ReturnDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
