package datamodel

import "gorm.io/gorm"

type Robber struct {
	*gorm.Model
	GameID uint
	Q      int
	R      int
	Status RobberStatus
}

type RobberStatus string

const (
	RS_IDLE = "IDLE"
	RS_MOVE = "MOVE"
	RS_ROB  = "ROB"
)
