package database

import (
	"github.com/wilburhimself/todo_go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB(dbPath string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// Run migrations
	err = db.AutoMigrate(&models.Todo{}, &models.User{})
	if err != nil {
		return err
	}

	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
