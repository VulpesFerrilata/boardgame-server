package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafePlayerRepository interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Players, error)
}

type PlayerRepository interface {
	SafePlayerRepository
	InsertOrUpdate(ctx context.Context, player *model.Player) error
	Delete(ctx context.Context, player *model.Player) error
}
