package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeDiceRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Dices, error)
}

type DiceRepository interface {
	SafeDiceRepository
	InsertOrUpdate(ctx context.Context, dice *model.Dice) error
}
