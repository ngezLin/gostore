package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CustomerID uint    `gorm:"not null"`
	Customer   User    `gorm:"foreignKey:CustomerID"`
	Status     string  `gorm:"type:ENUM('active', 'checked_out');default:'active'"`

	CartItems  []CartItem `gorm:"foreignKey:CartID"`
}
