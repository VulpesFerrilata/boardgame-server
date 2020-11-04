package datamodel

import "gorm.io/gorm"

type Road struct {
	*gorm.Model
	GameID   uint
	PlayerID uint
	Q        int
	R        int
	Location RoadLocation
}

type RoadLocation string

const (
	RL_TOP_LEFT RoadLocation = "TOP_LEFT"
	RL_MID_LEFT RoadLocation = "MID_LEFT"
	RL_BOT_LEFT RoadLocation = "BOT_LEFT"
)
