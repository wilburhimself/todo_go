package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Done  bool   `gorm:"not null, default:false"`
}
