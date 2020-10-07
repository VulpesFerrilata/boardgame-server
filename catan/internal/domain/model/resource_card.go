package model

import "gorm.io/gorm"

type ResourceCard struct {
	*gorm.Model
	GameID       uint
	PlayerID     uint
	ResourceType ResourceType
}

type ResourceType string

const (
	RT_BRICK ResourceType = "BRICK"
	RT_GRAIN ResourceType = "GRAIN"
	RT_SHEEP ResourceType = "SHEEP"
	RT_STONE ResourceType = "STONE"
	RT_WOOD  ResourceType = "WOOD"
)
