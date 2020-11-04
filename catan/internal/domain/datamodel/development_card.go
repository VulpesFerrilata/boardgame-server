package datamodel

import "gorm.io/gorm"

type DevelopmentCard struct {
	*gorm.Model
	GameID   uint
	PlayerID *uint
	Type     DevelopmentType
}

type DevelopmentType string

const (
	DT_KNIGHT         DevelopmentType = "KNIGHT"
	DT_MONOPOLY       DevelopmentType = "MONOPOLY"
	DT_ROAD_BUILDING  DevelopmentType = "ROAD_BUILDING"
	DT_YEAR_OF_PLENTY DevelopmentType = "YEAR_OF_PLENTY"
	DT_VICTORY_POINTS DevelopmentType = "VICTORY_POINTS"
)
