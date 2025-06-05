package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null"`
	Cart      Cart    `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`

	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`

	Quantity  int     `gorm:"not null"`
	SubTotal  float64 `gorm:"not null"` 
}
