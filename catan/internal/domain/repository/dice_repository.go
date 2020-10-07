package repository

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
)

type DiceRepository interface {
	FindByGameId(ctx context.Context, gameId uint) ([]*model.Dice, error)
	Insert(ctx context.Context, dices ...*model.Dice) error
	Save(ctx context.Context, dices ...*model.Dice) error
}
