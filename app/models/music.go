package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title  string `gorm:"size:200;not null"`
	Artist string `gorm:"size:100;not null"`
}
