package model

import "gorm.io/gorm"

type Robber struct {
	gorm.Model
	GameID uint
	Q      int
	R      int
}
