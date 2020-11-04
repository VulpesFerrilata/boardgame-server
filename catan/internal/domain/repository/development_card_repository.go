package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeDevelopmentCardRepository interface {
	FindByGameIdByPlayerId(ctx context.Context, gameId uint, playerId *uint) (model.DevelopmentCards, error)
}

type DevelopmentCardRepository interface {
	SafeDevelopmentCardRepository
	InsertOrUpdate(ctx context.Context, developmentCard *model.DevelopmentCard) error
}
