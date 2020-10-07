package model

import "gorm.io/gorm"

type Achievement struct {
	*gorm.Model
	GameID      uint
	PlayerID    uint
	Type        AchievementType
	BonusPoints int
}

type AchievementType string

const (
	AT_LONGEST_ROAD = "LONGEST_ROAD"
	AT_LARGEST_ARMY = "LARGEST_ARMY"
)
