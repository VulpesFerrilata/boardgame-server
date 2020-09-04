package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"type:"varchar(20);unique_index"`
	HashPassword []byte
}
