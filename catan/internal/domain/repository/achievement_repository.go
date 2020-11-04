package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeAchievementRepository interface {
	FindByGameIdByPlayerId(ctx context.Context, gameId uint, playerId *uint) (model.Achievements, error)
}

type AchievementRepository interface {
	SafeAchievementRepository
	InsertOrUpdate(ctx context.Context, achievement *model.Achievement) error
}
