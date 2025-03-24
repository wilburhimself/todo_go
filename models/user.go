package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:255;not null;unique"`
	Email     string `gorm:"size:255;not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Todos []Todo `gorm:"foreignKey:UserID"`
}

// BeforeSave will hash the password before saving the user
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Remove this hook as we handle hashing in SetPassword
	return nil
}

// SetPassword hashes the provided password and sets it to the user's Password field
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the provided password with the stored hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
