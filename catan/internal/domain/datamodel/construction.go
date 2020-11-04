package datamodel

import "gorm.io/gorm"

type Construction struct {
	*gorm.Model
	GameID           uint
	PlayerID         uint
	Q                int
	R                int
	Location         ConstructionLocation
	IsUpgradedCastle bool
}

type ConstructionLocation string

const (
	CL_TOP ConstructionLocation = "TOP"
	CL_BOT ConstructionLocation = "BOT"
)
