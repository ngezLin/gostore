package models

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}
