package datamodel

import "gorm.io/gorm"

type Game struct {
	*gorm.Model
	Host         uint
	PlayerInTurn uint
	Turn         int
	Status       GameStatus
}

type GameStatus string

const (
	GS_WAITING  GameStatus = "WAITING"
	GS_STARTED  GameStatus = "STARTED"
	GS_FINISHED GameStatus = "FINISHED"
)
