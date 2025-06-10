package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint
	Transaction   Transaction `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE"`

	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`

	Quantity int
	SubTotal float64
}
