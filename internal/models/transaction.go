package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CustomerID uint    `gorm:"not null"` // FK to users.id
	Customer   User    `gorm:"foreignKey:CustomerID"`

	CourierID  *uint   // Nullable — assigned later
	Courier    *User   `gorm:"foreignKey:CourierID"`

	TotalAmount float64 `gorm:"not null"`

	Status string `gorm:"type:ENUM('processing', 'delivering', 'arrived');default:'processing'"`

	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID"`
}
