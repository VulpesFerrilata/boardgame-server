package repository

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
)

type ReadOnlyDiceRepository interface {
	FindByGameId(ctx context.Context, gameId uint) ([]*model.Dice, error)
}

type DiceRepository interface {
	ReadOnlyDiceRepository
	Insert(ctx context.Context, dices ...*model.Dice) error
	Save(ctx context.Context, dices ...*model.Dice) error
}
