package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CustomerID uint   `gorm:"not null"`
	Customer   User   `gorm:"foreignKey:CustomerID"`

	CourierID *uint
	Courier   *User `gorm:"foreignKey:CourierID"`

	TotalAmount float64 `gorm:"not null;default:0"`
	Status      string  `gorm:"type:ENUM('cart', 'processing', 'delivering', 'arrived');default:'cart'"`

	Items []TransactionItem `gorm:"foreignKey:TransactionID"`
}
