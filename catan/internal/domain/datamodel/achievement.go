package datamodel

import "gorm.io/gorm"

type Achievement struct {
	*gorm.Model
	GameID      uint
	PlayerID    *uint
	Type        AchievementType
	BonusPoints int
}

type AchievementType string

const (
	AT_LONGEST_ROAD AchievementType = "LONGEST_ROAD"
	AT_LARGEST_ARMY AchievementType = "LARGEST_ARMY"
)
