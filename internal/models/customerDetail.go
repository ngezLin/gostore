package models

import "gorm.io/gorm"

type CustomerDetail struct {
	gorm.Model
	UserID      uint    `gorm:"unique;not null"` 
	User        User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	UserBalance float64 `gorm:"not null" json:"user_balance"`
}