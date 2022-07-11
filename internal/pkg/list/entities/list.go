package entities

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title       *string `gorm:"not null"`
	Description *string `gorm:"not null"`
	Priority    *int    `gorm:"default:0"`
}
