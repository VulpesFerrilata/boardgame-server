package model

import (
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	Jti string `gorm:"unique_index"`
}
