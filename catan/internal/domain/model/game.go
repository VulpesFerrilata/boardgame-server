package model

import "gorm.io/gorm"

type Game struct {
	*gorm.Model
	PlayerInTurn uint
	Turn         int
	Status       GameStatus
}

type GameStatus string

const (
	GS_WAITED   = "WAITED"
	GS_STARTED  = "STARTED"
	GS_FINISHED = "FINISHED"
)
