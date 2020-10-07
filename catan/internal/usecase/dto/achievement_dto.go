package dto

import "github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"

type AchievementDTO struct {
	ID          int
	Type        model.AchievementType
	BonusPoints int
}
