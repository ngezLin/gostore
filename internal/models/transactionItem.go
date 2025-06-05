package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint      `gorm:"not null"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE"`

	ProductID     uint      `gorm:"not null"`
	Product       Product   `gorm:"foreignKey:ProductID"`

	Quantity      int       `gorm:"not null"`
	SubTotal      float64   `gorm:"not null"`
}

