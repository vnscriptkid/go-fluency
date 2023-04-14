package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID uint `gorm:"foreignKey:UserRefer"`
	Street string
	City   string
	State  string
	Zip    string
}
