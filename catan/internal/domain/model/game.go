package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	CurrentTurnOrder  int
	CurrentTurn       int
	IsRolledDices     bool
	IsPlayedDevCard   bool
	IsUsingMonopoly   bool
	FreeRoads         int
	FreeResourceCards int
}
