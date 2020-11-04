package datamodel

import "gorm.io/gorm"

type Dice struct {
	*gorm.Model
	GameID   uint
	Number   int
	IsRolled bool
}
