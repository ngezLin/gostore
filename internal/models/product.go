package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name              string          `json:"name"`
	Price             float64         `json:"price"`
	Stock             int             `json:"stock"`
	ProductCategoryID uint            `json:"product_category_id"`
	Category          ProductCategory `gorm:"foreignKey:ProductCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
