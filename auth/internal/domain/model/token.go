package model

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Jti string `gorm:"unique_index"`
}
